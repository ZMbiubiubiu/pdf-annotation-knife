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
		Top:    200,
		Right:  200,
		Bottom: 100,
	}
	squareAnnot.Width = 10
	squareAnnot.StrikeColor = &color.RGBA{255, 0, 0, 60}
	squareAnnot.FillColor = &color.RGBA{0, 255, 0, 80}
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
				X: 105,
				Y: 105,
			},
			{
				X: 110,
				Y: 100,
			},
			{
				X: 115,
				Y: 95,
			},
			{
				X: 120,
				Y: 100,
			},
		},
	}
	inkAnnot.Rect = Rect{
		Left:   0,
		Bottom: 0,
		Top:    200,
		Right:  200,
	}
	inkAnnot.Width = 4
	inkAnnot.StrikeColor = &color.RGBA{255, 0, 255, 120}
	// inkAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
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

func TestAddFreeTextAnnotation(t *testing.T) {
	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_freetext.pdf")
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

	var freeTextAnnot = NewFreeTextAnnotation(page)
	freeTextAnnot.Rect = Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	}
	freeTextAnnot.Width = 2
	freeTextAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	freeTextAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	freeTextAnnot.Contents = "Hello, World!"
	freeTextAnnot.FontColor = &color.RGBA{255, 0, 0, 255}
	freeTextAnnot.GenerateAppearance()
	err = freeTextAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save freetext document failed: %v", err)
	}
}

func TestAddCircleAnnotation(t *testing.T) {
	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_circle.pdf")
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

	var circleAnnot = NewCircleAnnotation(page)
	circleAnnot.Rect = Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	}
	circleAnnot.Width = 2
	circleAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	circleAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	circleAnnot.GenerateAppearance()
	err = circleAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save circle document failed: %v", err)
	}
}

func TestAddHighlightAnnotation(t *testing.T) {
	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_highlight.pdf")
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

	var highlightAnnot = NewHighlightAnnotation(page)
	highlightAnnot.Rect = Rect{
		Left:   100,
		Top:    400,
		Right:  200,
		Bottom: 100,
	}
	highlightAnnot.Width = 2
	highlightAnnot.StrikeColor = &color.RGBA{255, 0, 0, 120}
	// highlightAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	highlightAnnot.QuadPoints = []QuadPoint{
		{
			LeftTopX:     100,
			LeftTopY:     200,
			RightTopX:    200,
			RightTopY:    200,
			RightBottomX: 200,
			RightBottomY: 100,
			LeftBottomX:  100,
			LeftBottomY:  100,
		},
		{
			LeftTopX:     100,
			LeftTopY:     400,
			RightTopX:    200,
			RightTopY:    400,
			RightBottomX: 200,
			RightBottomY: 300,
			LeftBottomX:  100,
			LeftBottomY:  300,
		},
	}
	highlightAnnot.GenerateAppearance()
	err = highlightAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save highlight document failed: %v", err)
	}
}

func TestAddUnderlineAnnotation(t *testing.T) {
	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_underline.pdf")
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

	var underlineAnnot = NewUnderlineAnnotation(page)
	underlineAnnot.Rect = Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	}
	underlineAnnot.Width = 2
	underlineAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	underlineAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	underlineAnnot.QuadPoints = []QuadPoint{
		{
			LeftTopX:     100,
			LeftTopY:     200,
			RightTopX:    200,
			RightTopY:    200,
			RightBottomX: 200,
			RightBottomY: 100,
			LeftBottomX:  100,
			LeftBottomY:  100,
		},
	}
	underlineAnnot.GenerateAppearance()
	err = underlineAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save underline document failed: %v", err)
	}
}

func TestAddStrikeoutAnnotation(t *testing.T) {
	inputFile := "simple.pdf"
	outputFile := strings.ReplaceAll(inputFile, ".pdf", "_add_strikeout.pdf")
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

	var strikeoutAnnot = NewStrikeoutAnnotation(page)
	strikeoutAnnot.Rect = Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	}
	strikeoutAnnot.Width = 2
	strikeoutAnnot.StrikeColor = &color.RGBA{255, 0, 0, 255}
	strikeoutAnnot.FillColor = &color.RGBA{0, 255, 0, 255}
	strikeoutAnnot.QuadPoints = []QuadPoint{
		{
			LeftTopX:     100,
			LeftTopY:     200,
			RightTopX:    200,
			RightTopY:    200,
			RightBottomX: 200,
			RightBottomY: 100,
			LeftBottomX:  100,
			LeftBottomY:  100,
		},
	}
	strikeoutAnnot.GenerateAppearance()
	err = strikeoutAnnot.AddAnnotationToPage(context.Background(), instance)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
		Document: docRes.Document,
		FilePath: &outputFile,
	})
	if err != nil {
		t.Fatalf("save strikeout document failed: %v", err)
	}
}
