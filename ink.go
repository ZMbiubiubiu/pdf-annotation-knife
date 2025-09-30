package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type InkAnnotation struct {
	BaseAnnotation
	Points [][]Point
}

func NewInkAnnotation(page requests.Page) *InkAnnotation {
	return &InkAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page:    page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_INK,
		},
	}
}

func (i *InkAnnotation) GenerateAppearance() error {
	// todo generate ink appearance
	i.AP = ""
	return nil
}

func (i *InkAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	err := i.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}

	// insert points
	for _, points := range i.Points {
		p := convertPointToPdfiumFormat(points)
		_, err = instance.FPDFAnnot_AddInkStroke(&requests.FPDFAnnot_AddInkStroke{
			Annotation: i.Annotation,
			Points:     p,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: i.Annotation,
	})
	if err != nil {
		return err
	}
	return nil
}
