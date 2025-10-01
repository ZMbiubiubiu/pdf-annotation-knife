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
	// generate highlight appearance
	h.AP = strings.Join([]string{
		h.GetColorAP(),
		h.GetPDFOpacityAP(),
		h.pointsCallback(h.QuadPoints),
		"f ",
	}, "\n")

	return nil
}

func (h *HighlightAnnotation) pointsCallback(quadPoints []QuadPoint) string {
	var pointsAP string
	for i := 0; i < len(quadPoints); i++ {
		pointsAP += fmt.Sprintf("%.3f %.3f m ", quadPoints[i].LeftTopX, quadPoints[i].LeftTopY)
		pointsAP += fmt.Sprintf("%.3f %.3f l ", quadPoints[i].RightTopX, quadPoints[i].RightTopY)
		pointsAP += fmt.Sprintf("%.3f %.3f l ", quadPoints[i].RightBottomX, quadPoints[i].RightBottomY)
		pointsAP += fmt.Sprintf("%.3f %.3f l ", quadPoints[i].LeftBottomX, quadPoints[i].LeftBottomY)
		pointsAP += "f\n"
	}
	return pointsAP
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
