// 矩形
package annotation

import (
	"context"
	"fmt"
	"strings"

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
			Page:    page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_SQUARE,
			NM:      GenerateUUID(),
		},
	}
}

func (s *SquareAnnotation) GenerateAppearance() error {
	// generate square appearance
	s.AP = strings.Join([]string{
		s.GetWidthAP(),
		s.GetColorAP(),
		s.GetPDFOpacityAP(),
		s.pointsCallback(),
	}, "\n")

	return nil
}

func (s *SquareAnnotation) pointsCallback() string {
	x := s.Rect.Left + float32(s.Width)/2
	y := s.Rect.Bottom + float32(s.Width)/2
	width := s.Rect.Right - s.Rect.Left - float32(s.Width)
	height := s.Rect.Top - s.Rect.Bottom - float32(s.Width)
	var ap = fmt.Sprintf("%.3f %.3f %.3f %.3f re\n", x, y, width, height)
	if s.FillColor != nil {
		ap += "B\n"
	}
	if s.StrikeColor != nil {
		ap += "S\n"
	}
	return ap
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
