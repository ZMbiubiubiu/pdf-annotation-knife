// 下划线
package annotation

import (
	"context"
	"fmt"
	"strings"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type UnderlineAnnotation struct {
	BaseAnnotation
	QuadPoints []QuadPoint
}

func NewUnderlineAnnotation() *UnderlineAnnotation {
	return &UnderlineAnnotation{
		BaseAnnotation: BaseAnnotation{
			Subtype: enums.FPDF_ANNOT_SUBTYPE_UNDERLINE,
			NM:      GenerateUUID(),
		},
	}
}

func (u *UnderlineAnnotation) GenerateAppearance() error {
	// generate underline appearance
	u.ap = strings.Join([]string{
		// "[] 0 d 0.571 w",
		"1 w",
		u.GetColorAP(),
		u.GetPDFOpacityAP(),
		u.pointsCallback(),
	}, "\n")
	return nil
}

func (u *UnderlineAnnotation) pointsCallback() string {
	var ap string
	for i := 0; i < len(u.QuadPoints); i++ {
		ap += fmt.Sprintf("%.3f %.3f m ", u.QuadPoints[i].LeftBottomX, u.QuadPoints[i].LeftBottomY+1.3)
		ap += fmt.Sprintf("%.3f %.3f l ", u.QuadPoints[i].RightBottomX, u.QuadPoints[i].RightBottomY+1.3)
		ap += "S\n"
	}
	return ap
}

func (u *UnderlineAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// create annotation
	err := u.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
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
