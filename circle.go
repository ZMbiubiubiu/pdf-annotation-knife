// 圆形
package annotation

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

var (
	// Circles are approximated by Bézier curves with four segments since
	// there is no circle primitive in the PDF specification. For the control
	// points distance, see https://stackoverflow.com/a/27863181.
	controlPointsDistance = float32((4.0 / 3.0) * math.Tan(math.Pi/(2.0*4)))
)

type CircleAnnotation struct {
	BaseAnnotation
}

func NewCircleAnnotation() *CircleAnnotation {
	return &CircleAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_CIRCLE,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
	}
}

func (c *CircleAnnotation) SetFillColor(color Color) {
	c.fillColor = &color
}

func (c *CircleAnnotation) GenerateAppearance() error {
	// generate circle appearance
	c.ap = strings.Join([]string{
		c.GetWidthAP(),
		c.GetColorAP(),
		c.GetPDFOpacityAP(),
		c.pointsCallback(),
	}, "\n")

	return nil
}

func (c *CircleAnnotation) pointsCallback() string {
	var ap string
	x0 := c.rect.Left + float32(c.width)/2
	y0 := c.rect.Top - float32(c.width)/2
	x1 := c.rect.Right - float32(c.width)/2
	y1 := c.rect.Bottom + float32(c.width)/2

	xMid := x0 + (x1-x0)/2
	yMid := y0 + (y1-y0)/2
	xOffset := ((x1 - x0) / 2) * controlPointsDistance
	yOffset := ((y1 - y0) / 2) * controlPointsDistance
	ap = strings.Join([]string{
		fmt.Sprintf("%.3f %.3f m", xMid, y1),
		fmt.Sprintf("%.3f %.3f %.3f %.3f %.3f %.3f c", xMid+xOffset, y1, x1, yMid+yOffset, x1, yMid),
		fmt.Sprintf("%.3f %.3f %.3f %.3f %.3f %.3f c", x1, yMid-yOffset, xMid+xOffset, y0, xMid, y0),
		fmt.Sprintf("%.3f %.3f %.3f %.3f %.3f %.3f c", xMid-xOffset, y0, x0, yMid-yOffset, x0, yMid),
		fmt.Sprintf("%.3f %.3f %.3f %.3f %.3f %.3f c", x0, yMid+yOffset, xMid-xOffset, y1, xMid, y1),
		"h\n",
	}, "\n")
	if c.fillColor != nil {
		ap += "B\n"
	}
	if c.strikeColor != nil {
		ap += "S\n"
	}

	return ap
}

func (c *CircleAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// create annotation
	err := c.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: c.annot,
	})
	if err != nil {
		return err
	}
	return nil
}
