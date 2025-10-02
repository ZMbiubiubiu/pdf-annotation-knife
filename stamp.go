package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type StampAnnotation struct {
	BaseAnnotation
}

func NewStampAnnotation(page requests.Page) *StampAnnotation {
	return &StampAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_STAMP,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
	}
}

func (s *StampAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// create annotation
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// insert object

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: s.annot,
	})
	if err != nil {
		return err
	}
	return nil
}
