package annotation

import (
	"context"
	"log"
	"os"
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
	outputFile := "data/simple_square.pdf"
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

	var squareAnnot = NewSquareAnnotation()
	squareAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
	squareAnnot.Width = 10
	squareAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
	squareAnnot.SetFillColor(Color{R: 0, G: 255, B: 0})
	squareAnnot.SetOpacity(120)
	squareAnnot.GenerateAppearance()
	err = squareAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	outputFile := "data/simple_ink.pdf"
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

	var inkAnnot = NewInkAnnotation()
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
	inkAnnot.SetRect(Rect{
		Left:   0,
		Bottom: 0,
		Top:    200,
		Right:  200,
	})
	inkAnnot.Width = 4
	inkAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 255})
	inkAnnot.SetOpacity(120)
	inkAnnot.GenerateAppearance()
	err = inkAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	outputFile := "data/simple_freetext.pdf"
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

	var freeTextAnnot = NewFreeTextAnnotation()
	freeTextAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
	freeTextAnnot.Width = 2
	freeTextAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
	freeTextAnnot.SetOpacity(120)
	freeTextAnnot.Contents = "Hello, World!"
	freeTextAnnot.FontColor = &Color{R: 255, G: 0, B: 0}
	freeTextAnnot.GenerateAppearance()
	err = freeTextAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	outputFile := "data/simple_circle.pdf"
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

	var circleAnnot = NewCircleAnnotation()
	circleAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
	circleAnnot.Width = 2
	circleAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
	circleAnnot.SetFillColor(Color{R: 0, G: 255, B: 0})
	circleAnnot.SetOpacity(120)
	circleAnnot.GenerateAppearance()
	err = circleAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	outputFile := "data/simple_highlight.pdf"
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

	var highlightAnnot = NewHighlightAnnotation()
	highlightAnnot.SetRect(Rect{
		Left:   0,
		Top:    400,
		Right:  200,
		Bottom: 0,
	})
	highlightAnnot.SetOpacity(120)
	highlightAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
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
	err = highlightAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	outputFile := "data/simple_underline.pdf"
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

	var underlineAnnot = NewUnderlineAnnotation()
	underlineAnnot.SetRect(Rect{
		Left:   0,
		Top:    410,
		Right:  210,
		Bottom: 0,
	})
	underlineAnnot.SetStrikeColor(Color{R: 0, G: 255, B: 0})
	underlineAnnot.SetOpacity(120)
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
	underlineAnnot.GenerateAppearance()
	err = underlineAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	outputFile := "data/simple_strikeout.pdf"
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

	var strikeoutAnnot = NewStrikeoutAnnotation()
	strikeoutAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
	strikeoutAnnot.Width = 2
	strikeoutAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
	strikeoutAnnot.SetOpacity(120)
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
	err = strikeoutAnnot.AddAnnotationToPage(context.Background(), instance, page)
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

func TestAddStampAnnotation(t *testing.T) {
	inputFile := "simple.pdf"

	t.Run("add path to stamp annot", func(t *testing.T) {

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

		var stampAnnot = NewStampAnnotation()
		stampAnnot.SetRect(Rect{
			Left:   0,
			Bottom: 0,
			Top:    200,
			Right:  200,
		})

		outputFile := "data/simple_stamp_path.pdf"
		os.Remove(outputFile)
		stampAnnot.SetPathObject([][]Point{
			{
				{X: 100, Y: 100},
				{X: 102, Y: 98},
				{X: 104, Y: 102},
				{X: 106, Y: 100},
				{X: 108, Y: 98},
				{X: 110, Y: 100},
				{X: 112, Y: 102},
				{X: 114, Y: 100},
				{X: 116, Y: 98},
				{X: 118, Y: 100},
				{X: 120, Y: 102},
			},
			{
				{X: 50, Y: 100},
				{X: 52, Y: 98},
				{X: 54, Y: 102},
				{X: 56, Y: 100},
				{X: 58, Y: 98},
				{X: 60, Y: 100},
				{X: 62, Y: 102},
				{X: 64, Y: 100},
				{X: 66, Y: 98},
				{X: 68, Y: 100},
				{X: 70, Y: 102},
			},
		}, 2, Color{R: 255, G: 0, B: 255}, 120)
		err = stampAnnot.AddAnnotationToPage(context.Background(), instance, page)
		if err != nil {
			t.Fatal(err)
		}

		_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
			Document: docRes.Document,
			FilePath: &outputFile,
		})
		if err != nil {
			t.Fatalf("save stamp document failed: %v", err)
		}
	})

	t.Run("add text to stamp annot", func(t *testing.T) {
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

		_ = page 
		// outputFile := "data/simple_text_stamp.pdf"
		// os.Remove(outputFile)
	})

	t.Run("add img to stamp annot", func(t *testing.T) {
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

		var stampAnnot = NewStampAnnotation()
		stampAnnot.SetRect(Rect{
			Left:   100,
			Bottom: 100,
			Top:    146,
			Right:  168,
		})

		outputFile := "data/simple_stamp_img.pdf"
		os.Remove(outputFile)

		// set image object
		stampAnnot.SetImgObject("jpeg", docRes.Document, "simple.jpeg")

		// add annotation to page
		err = stampAnnot.AddAnnotationToPage(context.Background(), instance, page)
		if err != nil {
			t.Fatal(err)
		}

		_, err = instance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
			Document: docRes.Document,
			FilePath: &outputFile,
		})
		if err != nil {
			t.Fatalf("save img stamp document failed: %v", err)
		}
	})
}
