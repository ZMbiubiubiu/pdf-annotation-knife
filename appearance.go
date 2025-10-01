package annotation

import (
	"fmt"
	"image/color"
	"strings"
)

func (b *BaseAnnotation) getColorAP(color *color.RGBA, isFill bool) string {

	if color == nil {
		return ""
	}
	
	var op string
	if color.R == color.G && color.G == color.B {
		op = "G" // Set gray level for stroking (outline) operations.
		if isFill {
			op = "g" // Set gray level for non-stroking (fill) operations.
		}
		return fmt.Sprintf("%.3f %s", float32(color.R)/255.0, op)
	}
	op = "RG" // Set RGB color for stroking (outline) operations.
	if isFill {
		op = "rg" //Set RGB color for non-stroking (fill) operations.
	}
	return fmt.Sprintf("%.3f %.3f %.3f %s", float32(color.R)/255.0, float32(color.G)/255.0, float32(color.B)/255.0, op)
}

func (b *BaseAnnotation) GetWidthAP() string {
	return fmt.Sprintf("%.3f w", float32(b.Width))
}

func (b *BaseAnnotation) GetColorAP() string {
	return strings.Join([]string{
		b.getColorAP(b.fillColor, true),
		b.getColorAP(b.strikeColor, false),
	}, "\n")
}

func (b *BaseAnnotation) GetPDFOpacityAP() string {
	return "/GS gs"
}
