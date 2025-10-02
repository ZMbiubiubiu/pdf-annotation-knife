// 高亮
package annotation

import (
	"context"
	"fmt"
	"strings"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

var (
	DefaultHighlightColor = Color{R: 255, G: 255, B: 0} // Default color is yellow in Acrobat Reader
)

type HighlightAnnotation struct {
	BaseAnnotation
	QuadPoints []QuadPoint
}

func NewHighlightAnnotation() *HighlightAnnotation {
	return &HighlightAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_HIGHLIGHT,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
	}
}

func (h *HighlightAnnotation) GenerateAppearance() error {
	// generate highlight appearance
	h.ap = strings.Join([]string{
		h.GetPDFOpacityAP(),
		h.GetColorAP(),
		h.pointsCallback(),
	}, "\n")

	return nil
}

// SetStrikeColor sets the color of the highlight annotation.
// highlight is different from underline&strikeout
// it should set fill color instead of stroke color
func (h *HighlightAnnotation) SetStrikeColor(c Color)  {

	h.fillColor = &c
}

func (h *HighlightAnnotation) pointsCallback() string {
	var pointsAP string
	for i := 0; i < len(h.QuadPoints); i++ {
		pointsAP += fmt.Sprintf("%.3f %.3f m ", h.QuadPoints[i].LeftTopX, h.QuadPoints[i].LeftTopY)
		pointsAP += fmt.Sprintf("%.3f %.3f l ", h.QuadPoints[i].RightTopX, h.QuadPoints[i].RightTopY)
		pointsAP += fmt.Sprintf("%.3f %.3f l ", h.QuadPoints[i].RightBottomX, h.QuadPoints[i].RightBottomY)
		pointsAP += fmt.Sprintf("%.3f %.3f l ", h.QuadPoints[i].LeftBottomX, h.QuadPoints[i].LeftBottomY)
		pointsAP += "h f\n"
	}
	return pointsAP
}

func (h *HighlightAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	if h.strikeColor == nil {
		h.strikeColor = &DefaultHighlightColor
	}
	// create annotation
	err := h.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// insert quad points
	quadPoints := convertQuadPointToPdfiumFormat(h.QuadPoints)
	for _, points := range quadPoints {
		_, err = instance.FPDFAnnot_AppendAttachmentPoints(&requests.FPDFAnnot_AppendAttachmentPoints{
			Annotation:       h.annot,
			AttachmentPoints: points,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: h.annot,
	})
	if err != nil {
		return err
	}

	return nil
}
