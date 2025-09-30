// 下划线
package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type UnderlineAnnotation struct {
	BaseAnnotation
	QuadPoints []QuadPoint
}

func NewUnderlineAnnotation(page requests.Page) *UnderlineAnnotation {
	return &UnderlineAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page:    page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_UNDERLINE,
		},
	}
}

func (u *UnderlineAnnotation) GenerateAppearance() error {
	// todo generate underline appearance
	u.AP = ""
	return nil
}

func (u *UnderlineAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	// create annotation
	err := u.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}

	// insert quad points
	quadPoints := convertQuadPointToPdfiumFormat(u.QuadPoints)
	for _, points := range quadPoints {
		_, err = instance.FPDFAnnot_AppendAttachmentPoints(&requests.FPDFAnnot_AppendAttachmentPoints{
			Annotation:       u.Annotation,
			AttachmentPoints: points,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: u.Annotation,
	})
	if err != nil {
		return err
	}
	return nil
}
