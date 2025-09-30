// 矩形
package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type SquareAnnotation struct {
	BaseAnnotation
}

func NewSquareAnnotation(page requests.Page) *SquareAnnotation {
	return &SquareAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page: page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_SQUARE,
		},
	}
}

func (s *SquareAnnotation) GenerateAppearance() error {
	// todo generate square appearance
	s.AP = ""
	return nil
}

func (s *SquareAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}

	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: s.Annotation,
	})
	if err != nil {
		return err
	}

	return nil
}
