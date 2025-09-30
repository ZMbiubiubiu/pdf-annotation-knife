// 高亮
package annotation

import (
	"context"
	"fmt"
	"image/color"
	"strings"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

var (
	DefaultHighlightColor = color.RGBA{255, 255, 0, 255} // Default color is yellow in Acrobat Reader
)

type HighlightAnnotation struct {
	BaseAnnotation
	QuadPoints []QuadPoint
}

func NewHighlightAnnotation(page requests.Page) *HighlightAnnotation {
	return &HighlightAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page:    page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_HIGHLIGHT,
			NM:      GenerateUUID(),
		},
	}
}

func (h *HighlightAnnotation) GenerateAppearance() error {
	// todo generate highlight appearance
	h.AP = strings.Join([]string{
		h.GetPDFColorAP(h.StrikeColor, false),
		h.GetPDFOpacityAP(),
		h.PointsCallback(h.QuadPoints),
		"f ",
	}, "\n")

	return nil
}

func (h *HighlightAnnotation) PointsCallback(quadPoints []QuadPoint) string {
	var points string
	for i := 0; i < len(quadPoints); i++ {
		points += fmt.Sprintf("%f %f m ", quadPoints[i].LeftTopX, quadPoints[i].LeftTopY)
		points += fmt.Sprintf("%f %f l ", quadPoints[i].RightTopX, quadPoints[i].RightTopY)
		points += fmt.Sprintf("%f %f l ", quadPoints[i].RightBottomX, quadPoints[i].RightBottomY)
		points += fmt.Sprintf("%f %f l ", quadPoints[i].LeftBottomX, quadPoints[i].LeftBottomY)
	}
	return points
}

func (h *HighlightAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	if h.StrikeColor == nil {
		h.StrikeColor = &DefaultHighlightColor
	}
	// create annotation
	err := h.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}

	// insert quad points
	quadPoints := convertQuadPointToPdfiumFormat(h.QuadPoints)
	for _, points := range quadPoints {
		_, err = instance.FPDFAnnot_AppendAttachmentPoints(&requests.FPDFAnnot_AppendAttachmentPoints{
			Annotation:       h.Annotation,
			AttachmentPoints: points,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: h.Annotation,
	})
	if err != nil {
		return err
	}

	return nil
}
