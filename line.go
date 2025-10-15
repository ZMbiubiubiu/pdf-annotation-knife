// Line 直线
package annotation

import (
	"context"
	"fmt"
	"strings"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type LineAnnotation struct {
	BaseAnnotation
	lineTo         [2]Point
	strikeLineCap  enums.FPDF_LINECAP
	strikeLineJoin enums.FPDF_LINEJOIN
}

func NewLineAnnotation() *LineAnnotation {
	return &LineAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_LINE,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
		strikeLineCap:  enums.FPDF_LINECAP_BUTT,
		strikeLineJoin: enums.FPDF_LINEJOIN_MITER,
	}
}

func (l *LineAnnotation) SetLineTo(beginX, beginY, endX, endY float32) {
	l.lineTo = [2]Point{
		{X: beginX, Y: beginY},
		{X: endX, Y: endY},
	}
}

// SetStrikeLineCap sets the line cap style for the stroke.
func (l *LineAnnotation) SetStrikeLineCap(lineCap enums.FPDF_LINECAP) {
	l.strikeLineCap = lineCap
}

// SetStrikeLineJoin sets the line join style for the stroke.
func (l *LineAnnotation) SetStrikeLineJoin(lineJoin enums.FPDF_LINEJOIN) {
	l.strikeLineJoin = lineJoin
}

func (l *LineAnnotation) GetLineStyleAP() string {
	return fmt.Sprintf("%d j %d J", l.strikeLineCap, l.strikeLineJoin)
}

// GenerateAppearance generates the appearance stream for the line annotation.
func (l *LineAnnotation) GenerateAppearance() error {
	l.ap = strings.Join([]string{
		l.GetColorAP(),
		l.GetWidthAP(),
		l.GetPDFOpacityAP(),
		l.GetLineStyleAP(),
		l.pointsCallback(),
	}, "\n")

	return nil
}

func (l *LineAnnotation) pointsCallback() string {
	var path string
	for i, point := range l.lineTo {
		op := "l"
		if i == 0 {
			op = "m"
		}
		path += fmt.Sprintf("%.3f %.3f %s ", point.X, point.Y, op)

	}
	path += "S "

	return path
}

func (l *LineAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// create annotation
	err := l.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: l.annot,
	})
	if err != nil {
		return err
	}
	return nil
}
