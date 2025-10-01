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
			Page: page,
			// 由于 enums 未定义，推测可能需要从 go-pdfium 包中导入对应枚举
			Subtype: enums.FPDF_ANNOT_SUBTYPE_STAMP,
			NM:      GenerateUUID(),
		},
	}
}

func (s *StampAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	// create annotation
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
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