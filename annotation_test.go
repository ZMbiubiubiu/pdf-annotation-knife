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

func TestAddLineAnnotation(t *testing.T) {
	inputFile := "simple.pdf"

	t.Run("simple normarl line", func(t *testing.T) {
		// output file
		outputFile := "data/simple_line.pdf"
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

		var lineAnnot = NewLineAnnotation()
		lineAnnot.SetRect(Rect{
			Left:   0,
			Top:    700,
			Right:  600,
			Bottom: 0,
		})
		lineAnnot.SetWidth(8)
		lineAnnot.SetStrikeColor(Color{R: 0, G: 120, B: 0})
		lineAnnot.SetOpacity(120)
		lineAnnot.SetLineTo(100, 200, 200, 100)
		lineAnnot.GenerateAppearance()

		err = lineAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	})

	t.Run("simple open arrow line", func(t *testing.T) {
		// output file
		outputFile := "data/simple_open_arrow_line.pdf"
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

		var lineAnnot = NewLineAnnotation()
		lineAnnot.SetRect(Rect{
			Left:   0,
			Top:    700,
			Right:  600,
			Bottom: 0,
		})
		lineAnnot.SetWidth(2)
		lineAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
		lineAnnot.SetOpacity(120)
		lineAnnot.SetLineTo(150.799, 603.666, 472.192, 652.451)
		lineAnnot.SetCustomAppearance(`1.0 0.0 0.0 RG
/GS gs 
18.000000 w
2 J 
150.799 603.666 m
472.192 652.451 l
441.962 609.090 m
501.853 656.954 l
430.457 684.889 l
S`)
		err = lineAnnot.AddAnnotationToPage(context.Background(), instance, page)
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
	})
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
	squareAnnot.SetWidth(10)
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
			{X: 52.691, Y: 777.755},
			{X: 50.555, Y: 779.18},
			{X: 49.843, Y: 779.18},
			{X: 49.487, Y: 779.536},
			{X: 49.131, Y: 779.536},
			{X: 48.419, Y: 779.536},
			{X: 48.063, Y: 779.892},
			{X: 47.351, Y: 779.892},
			{X: 46.639, Y: 779.892},
			{X: 45.927, Y: 779.892},
			{X: 44.859, Y: 779.892},
			{X: 43.791, Y: 779.892},
			{X: 42.723, Y: 779.892},
			{X: 41.654, Y: 779.536},
			{X: 40.586, Y: 779.536},
			{X: 39.518, Y: 779.18},
			{X: 38.45, Y: 778.824},
			{X: 37.382, Y: 778.824},
			{X: 36.314, Y: 778.468},
			{X: 35.246, Y: 778.111},
			{X: 34.178, Y: 777.755},
			{X: 33.466, Y: 777.399},
			{X: 32.398, Y: 776.687},
			{X: 31.33, Y: 776.331},
			{X: 30.618, Y: 775.975},
			{X: 29.55, Y: 775.263},
			{X: 28.838, Y: 774.906},
			{X: 28.482, Y: 774.194},
			{X: 27.77, Y: 773.482},
			{X: 27.058, Y: 772.414},
			{X: 26.702, Y: 771.701},
			{X: 26.346, Y: 770.989},
			{X: 25.99, Y: 770.277},
			{X: 25.634, Y: 769.209},
			{X: 25.634, Y: 768.496},
			{X: 25.634, Y: 767.428},
			{X: 25.277, Y: 766.716},
			{X: 25.634, Y: 766.004},
			{X: 25.634, Y: 764.935},
			{X: 25.99, Y: 764.223},
			{X: 26.346, Y: 763.511},
			{X: 26.702, Y: 762.442},
			{X: 27.414, Y: 761.374},
			{X: 28.126, Y: 760.662},
			{X: 28.838, Y: 759.594},
			{X: 29.55, Y: 758.881},
			{X: 30.618, Y: 758.169},
			{X: 32.042, Y: 757.101},
			{X: 33.11, Y: 756.388},
			{X: 34.534, Y: 755.676},
			{X: 35.958, Y: 754.964},
			{X: 37.738, Y: 754.608},
			{X: 39.518, Y: 754.252},
			{X: 41.298, Y: 753.54},
			{X: 43.079, Y: 753.183},
			{X: 45.215, Y: 753.183},
			{X: 48.775, Y: 752.471},
			{X: 50.911, Y: 752.471},
			{X: 53.047, Y: 752.471},
			{X: 54.827, Y: 752.471},
			{X: 56.963, Y: 752.471},
			{X: 59.099, Y: 752.827},
			{X: 61.236, Y: 752.827},
			{X: 63.372, Y: 753.183},
			{X: 65.508, Y: 753.54},
			{X: 67.644, Y: 754.252},
			{X: 69.78, Y: 754.608},
			{X: 71.56, Y: 755.32},
			{X: 73.34, Y: 756.032},
			{X: 74.764, Y: 756.388},
			{X: 76.188, Y: 757.101},
			{X: 77.613, Y: 757.813},
			{X: 78.681, Y: 758.169},
			{X: 80.105, Y: 758.881},
			{X: 81.173, Y: 759.594},
			{X: 81.885, Y: 760.306},
			{X: 82.597, Y: 761.018},
			{X: 83.309, Y: 761.73},
			{X: 84.021, Y: 762.086},
			{X: 84.733, Y: 762.799},
			{X: 85.089, Y: 763.511},
			{X: 85.801, Y: 764.579},
			{X: 86.157, Y: 765.291},
			{X: 86.157, Y: 766.36},
			{X: 86.513, Y: 767.072},
			{X: 86.869, Y: 767.784},
			{X: 86.869, Y: 768.853},
			{X: 86.869, Y: 769.921},
			{X: 86.513, Y: 770.633},
			{X: 86.513, Y: 771.701},
			{X: 85.801, Y: 772.414},
			{X: 85.445, Y: 773.126},
			{X: 84.733, Y: 773.482},
			{X: 84.021, Y: 774.194},
			{X: 82.953, Y: 774.55},
			{X: 81.529, Y: 775.263},
			{X: 80.461, Y: 775.619},
			{X: 79.037, Y: 775.619},
			{X: 77.613, Y: 775.975},
			{X: 75.832, Y: 775.975},
			{X: 74.408, Y: 776.331},
			{X: 72.984, Y: 776.331},
			{X: 71.204, Y: 776.331},
			{X: 69.424, Y: 776.331},
			{X: 67.644, Y: 776.331},
			{X: 65.508, Y: 775.975},
			{X: 63.728, Y: 775.619},
			{X: 62.304, Y: 775.619},
			{X: 60.524, Y: 775.263},
			{X: 58.743, Y: 774.906},
			{X: 57.319, Y: 774.55},
			{X: 55.895, Y: 774.194},
			{X: 54.827, Y: 773.838},
			{X: 53.759, Y: 773.838},
			{X: 52.335, Y: 773.482},
			{X: 51.623, Y: 773.126},
			{X: 50.911, Y: 772.77},
			{X: 50.199, Y: 772.414},
			{X: 49.487, Y: 772.414},
			{X: 48.775, Y: 771.701},
			{X: 47.707, Y: 771.345},
		},
	}
	inkAnnot.SetRect(Rect{
		Left:   0,
		Bottom: 700,
		Top:    800,
		Right:  100,
	})
	inkAnnot.SetWidth(4)
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
	freeTextAnnot.SetWidth(2)
	freeTextAnnot.SetOpacity(120)
	freeTextAnnot.Contents = "Hello, World!"
	freeTextAnnot.FontColor = Color{R: 255, G: 0, B: 0}
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
	circleAnnot.SetWidth(2)
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
	inputFile := "simple_text.pdf"
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
		Left:   239.084,
		Top:    446.565,
		Right:  932.784,
		Bottom: 387.359,
	})
	highlightAnnot.SetOpacity(255)
	highlightAnnot.SetStrikeColor(Color{R: 255, G: 255, B: 0})
	highlightAnnot.QuadPoints = []QuadPoint{
		{
			LeftTopX:     239.85,
			LeftTopY:     445.799,
			RightTopX:    932.018,
			RightTopY:    445.799,
			LeftBottomX:  239.85,
			LeftBottomY:  421.276,
			RightBottomX: 932.018,
			RightBottomY: 421.276,
		},
		{
			LeftTopX:     239.85,
			LeftTopY:     412.649,
			RightTopX:    854.006,
			RightTopY:    412.649,
			LeftBottomX:  239.85,
			LeftBottomY:  388.126,
			RightBottomX: 854.006,
			RightBottomY: 388.126,
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
	inputFile := "simple_text.pdf"
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
		Left:   239.084,
		Top:    446.565,
		Right:  932.784,
		Bottom: 387.359,
	})
	underlineAnnot.SetStrikeColor(Color{R: 0, G: 255, B: 0})
	underlineAnnot.SetOpacity(120)
	underlineAnnot.QuadPoints = []QuadPoint{
		{
			LeftTopX:     239.85,
			LeftTopY:     445.799,
			RightTopX:    932.018,
			RightTopY:    445.799,
			LeftBottomX:  239.85,
			LeftBottomY:  421.276,
			RightBottomX: 932.018,
			RightBottomY: 421.276,
		},
		{
			LeftTopX:     239.85,
			LeftTopY:     412.649,
			RightTopX:    854.006,
			RightTopY:    412.649,
			LeftBottomX:  239.85,
			LeftBottomY:  388.126,
			RightBottomX: 854.006,
			RightBottomY: 388.126,
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
	inputFile := "simple_text.pdf"
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
		Left:   239.084,
		Top:    446.565,
		Right:  932.784,
		Bottom: 387.359,
	})
	strikeoutAnnot.SetWidth(4)
	strikeoutAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
	strikeoutAnnot.SetOpacity(255)
	strikeoutAnnot.QuadPoints = []QuadPoint{
		{
			LeftTopX:     239.85,
			LeftTopY:     445.799,
			RightTopX:    932.018,
			RightTopY:    445.799,
			LeftBottomX:  239.85,
			LeftBottomY:  421.276,
			RightBottomX: 932.018,
			RightBottomY: 421.276,
		},
		{
			LeftTopX:     239.85,
			LeftTopY:     412.649,
			RightTopX:    854.006,
			RightTopY:    412.649,
			LeftBottomX:  239.85,
			LeftBottomY:  388.126,
			RightBottomX: 854.006,
			RightBottomY: 388.126,
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
