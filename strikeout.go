// 删除线
package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type StrikeoutAnnotation struct {
	BaseAnnotation
	QuadPoints []QuadPoint
}

func NewStrikeoutAnnotation(page requests.Page) *StrikeoutAnnotation {
	return &StrikeoutAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page:    page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_STRIKEOUT,
		},
	}
}

func (s *StrikeoutAnnotation) GenerateAppearance() error {
	// todo generate strikeout appearance
	s.AP = ""
	return nil
}

func (s *StrikeoutAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	// create annotation
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}

	// insert quad points
	quadPoints := convertQuadPointToPdfiumFormat(s.QuadPoints)
	for _, points := range quadPoints {
		_, err = instance.FPDFAnnot_AppendAttachmentPoints(&requests.FPDFAnnot_AppendAttachmentPoints{
			Annotation:       s.Annotation,
			AttachmentPoints: points,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: s.Annotation,
	})
	if err != nil {
		return err
	}

	return nil
}
