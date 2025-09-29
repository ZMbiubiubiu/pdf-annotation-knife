// 圆形
package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/requests"
)

type CircleAnnotation struct {
	BaseAnnotation
}

func NewCircleAnnotation(page requests.Page) *CircleAnnotation {
	return &CircleAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page: page,
		},
	}
}

func (c *CircleAnnotation) GenerateAppearance() error {
	// todo generate circle appearance
	c.AP = ""
	return nil
}

func (c *CircleAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	err := c.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}
	
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: c.Annotation,
	})
	if err != nil {
		return err
	}
	return nil
}


