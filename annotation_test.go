package annotation

import (
	"image/color"
	"log"
	"testing"
	"time"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/single_threaded"
)

// Be sure to close pools/instances when you're done with them.
var pool pdfium.Pool
var instance pdfium.Pdfium

func init() {
	// Init the PDFium library and return the instance to open documents.
	pool = single_threaded.Init(single_threaded.Config{})

	var err error
	instance, err = pool.GetInstance(time.Second * 30)
	if err != nil {
		log.Fatal(err)
	}
}

func TestAddAnnotation(t *testing.T) {
	// new square

	var squareAnnot = NewSquareAnnotation()
	squareAnnot.Rect = Rect{
		Left:   100,
		Top:    100,
		Right:  200,
		Bottom: 200,
	}
	squareAnnot.Width = 2
	squareAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	squareAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
}
