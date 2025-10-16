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
	LineStyle
	lineTo [2]Point
}

func NewLineAnnotation() *LineAnnotation {
	return &LineAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_LINE,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
		LineStyle: LineStyle{
			StrikeLineCap:  enums.FPDF_LINECAP_BUTT,
			StrikeLineJoin: enums.FPDF_LINEJOIN_MITER,
		},
	}
}

func (l *LineAnnotation) SetLineTo(beginX, beginY, endX, endY float32) {
	l.lineTo = [2]Point{
		{X: beginX, Y: beginY},
		{X: endX, Y: endY},
	}
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
