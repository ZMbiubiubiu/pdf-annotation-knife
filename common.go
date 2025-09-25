package annotation

import (
	"image/color"

	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

// annotation width
type Width float32

type Rect struct {
	Left, Bottom, Right, Top float32
}

// strike color
type StrikeColor color.RGBA

type BaseAnnotation struct {
	Page        requests.Page
	Subtype     enums.FPDF_ANNOTATION_SUBTYPE
	Rect        Rect
	Width       Width
	StrikeColor StrikeColor
}

func (b *BaseAnnotation) CreateAnnotation() error {
	return nil
}
