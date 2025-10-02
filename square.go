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

func NewSquareAnnotation() *SquareAnnotation {
	return &SquareAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_SQUARE,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
	}
}

func (s *SquareAnnotation) SetFillColor(c Color) {
	s.fillColor = &c
}

func (s *SquareAnnotation) GenerateAppearance() error {
	// generate square appearance
	s.ap = strings.Join([]string{
		s.GetWidthAP(),
		s.GetColorAP(),
		s.GetPDFOpacityAP(),
		s.pointsCallback(),
	}, "\n")

	return nil
}

func (s *SquareAnnotation) pointsCallback() string {
	x := s.rect.Left + float32(s.Width)/2
	y := s.rect.Bottom + float32(s.Width)/2
	width := s.rect.Right - s.rect.Left - float32(s.Width)
	height := s.rect.Top - s.rect.Bottom - float32(s.Width)
	var ap = fmt.Sprintf("%.3f %.3f %.3f %.3f re\n", x, y, width, height)
	if s.fillColor != nil {
		ap += "B\n"
	}
	if s.strikeColor != nil {
		ap += "S\n"
	}
	return ap
}

func (s *SquareAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: s.annot,
	})
	if err != nil {
		return err
	}

	return nil
}
