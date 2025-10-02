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

func NewStrikeoutAnnotation() *StrikeoutAnnotation {
	return &StrikeoutAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_STRIKEOUT,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
	}
}

func (s *StrikeoutAnnotation) GenerateAppearance() error {
	// todo generate strikeout appearance
	s.ap = strings.Join([]string{
		"[] 0 d 1 w",
		s.GetColorAP(),
		s.GetPDFOpacityAP(),
		s.pointsCallback(),
	}, "\n")
	return nil
}

func (s *StrikeoutAnnotation) pointsCallback() string {
	var pointsAp string
	for i := 0; i < len(s.QuadPoints); i++ {
		pointsAp += fmt.Sprintf("%.3f %.3f m ", (s.QuadPoints[i].LeftTopX+s.QuadPoints[i].LeftBottomX)/2, (s.QuadPoints[i].LeftTopY+s.QuadPoints[i].LeftBottomY)/2)
		pointsAp += fmt.Sprintf("%.3f %.3f l ", (s.QuadPoints[i].RightTopX+s.QuadPoints[i].RightBottomX)/2, (s.QuadPoints[i].RightTopY+s.QuadPoints[i].RightBottomY)/2)
		pointsAp += "S\n"
	}
	return pointsAp
}

func (s *StrikeoutAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// create annotation
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// insert quad points
	quadPoints := convertQuadPointToPdfiumFormat(s.QuadPoints)
	for _, points := range quadPoints {
		_, err = instance.FPDFAnnot_AppendAttachmentPoints(&requests.FPDFAnnot_AppendAttachmentPoints{
			Annotation:       s.annot,
			AttachmentPoints: points,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: s.annot,
	})
	if err != nil {
		return err
	}

	return nil
}
