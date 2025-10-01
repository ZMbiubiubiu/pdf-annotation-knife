// 删除线
package annotation

import (
	"context"
	"fmt"
	"strings"

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
			NM:      GenerateUUID(),
		},
	}
}

func (s *StrikeoutAnnotation) GenerateAppearance() error {
	// todo generate strikeout appearance
	s.AP = strings.Join([]string{
		"[] 0 d 1 w",
		s.GetColorAP(),
		s.GetPDFOpacityAP(),
		s.pointsCallback(),
	}, "\n")
	return nil
}

func (s *StrikeoutAnnotation) pointsCallback() string {
	var ap string
	for i := 0; i < len(s.QuadPoints); i++ {
		ap += fmt.Sprintf("%.3f %.3f m ", (s.QuadPoints[i].LeftTopX+s.QuadPoints[i].LeftBottomX)/2, (s.QuadPoints[i].LeftTopY+s.QuadPoints[i].LeftBottomY)/2)
		ap += fmt.Sprintf("%.3f %.3f l ", (s.QuadPoints[i].RightTopX+s.QuadPoints[i].RightBottomX)/2, (s.QuadPoints[i].RightTopY+s.QuadPoints[i].RightBottomY)/2)
		ap += "S\n"
	}
	return ap
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
