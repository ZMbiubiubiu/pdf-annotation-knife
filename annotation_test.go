package annotation

import (
	"context"
	"image/color"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/requests"
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

func TestAddSquareAnnotation(t *testing.T) {

	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_square.pdf")
	os.Remove(outputFile)
	docRes, err := instance.OpenDocument(&requests.OpenDocument{
		FilePath: &inputFile,
	})
	if err != nil {
		t.Fatalf("open document failed: %v", err)
	}

	page := requests.Page{
		ByIndex: &requests.PageByIndex{
			Document: docRes.Document,
			Index:    0,
		},
	}

	var squareAnnot = NewSquareAnnotation(page)
	squareAnnot.Rect = Rect{
		Left:   100,
		Top:    100,
		Right:  200,
		Bottom: 200,
	}
	squareAnnot.Width = 2
	squareAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	squareAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	squareAnnot.GenerateAppearance()
	err = squareAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save square document failed: %v", err)
	}
}

func TestAddInkAnnotation(t *testing.T) {

	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_ink.pdf")
	os.Remove(outputFile)
	docRes, err := instance.OpenDocument(&requests.OpenDocument{
		FilePath: &inputFile,
	})
	if err != nil {
		t.Fatalf("open document failed: %v", err)
	}

	page := requests.Page{
		ByIndex: &requests.PageByIndex{
			Document: docRes.Document,
			Index:    0,
		},
	}

	var inkAnnot = NewInkAnnotation(page)
	inkAnnot.Points = [][]Point{
		{
			{
				X: 100,
				Y: 100,
			},
			{
				X: 200,
				Y: 200,
			},
		},
	}
	inkAnnot.Rect = Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	}
	inkAnnot.Width = 2
	inkAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	inkAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	inkAnnot.GenerateAppearance()
	err = inkAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save ink document failed: %v", err)
	}
}
