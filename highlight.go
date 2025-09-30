// 高亮
package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
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
		},
	}
}

func (h *HighlightAnnotation) GenerateAppearance() error {
	// todo generate highlight appearance
	h.AP = ""
	return nil
}

func (h *HighlightAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
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
