package annotation

import (
	"context"
	"errors"
	"log"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/structs"
)

const (
	DefaultOpacity = 255 // default opacity is 100%
)

type Rect struct {
	Left, Bottom, Right, Top float32
}

type Point struct {
	X, Y float32
}

func convertPointToPdfiumFormat(points []Point) []structs.FPDF_FS_POINTF {
	pdfiumPoints := make([]structs.FPDF_FS_POINTF, 0, len(points))
	for _, point := range points {
		pdfiumPoints = append(pdfiumPoints, structs.FPDF_FS_POINTF{
			X: point.X,
			Y: point.Y,
		})
	}
	return pdfiumPoints
}

type QuadPoint struct {
	LeftTopX, LeftTopY         float32
	RightTopX, RightTopY       float32
	RightBottomX, RightBottomY float32
	LeftBottomX, LeftBottomY   float32
}

func convertQuadPointToPdfiumFormat(quadPoints []QuadPoint) []structs.FPDF_FS_QUADPOINTSF {
	pdfiumQuadPoints := make([]structs.FPDF_FS_QUADPOINTSF, 0, len(quadPoints))
	for _, quadPoint := range quadPoints {
		pdfiumQuadPoints = append(pdfiumQuadPoints, structs.FPDF_FS_QUADPOINTSF{
			X1: quadPoint.LeftTopX,
			Y1: quadPoint.LeftTopY,
			X2: quadPoint.RightTopX,
			Y2: quadPoint.RightTopY,
			X3: quadPoint.LeftBottomX,
			Y3: quadPoint.LeftBottomY,
			X4: quadPoint.RightBottomX,
			Y4: quadPoint.RightBottomY,
		})
	}
	return pdfiumQuadPoints
}

type BorderStyle struct {
	Width     float32
	Style     string
	DashArray []int
	Effect    string
	EffectInt int
}

type Color struct {
	R, G, B uint8
}

type BaseAnnotation struct {
	nm          string // annotation unique name
	title       string
	annot       references.FPDF_ANNOTATION
	subtype     enums.FPDF_ANNOTATION_SUBTYPE
	rect        Rect
	Width       float32
	opacity     uint8 // [0 -255]
	strikeColor *Color
	fillColor   *Color
	ap          string
}

// SetTitle sets the title of the annotation.
func (b *BaseAnnotation) SetTitle(title string) {
	b.title = title
}

// SetRect sets the rect of the annotation.
func (b *BaseAnnotation) SetRect(rect Rect) {
	b.rect = rect
}

// SetOpacity sets the opacity of the annotation.
func (b *BaseAnnotation) SetOpacity(opacity uint8) {
	b.opacity = opacity
}

// SetStrikeColor sets the strike color of the annotation.
func (b *BaseAnnotation) SetStrikeColor(c Color) {
	b.strikeColor = &c
}

func (b *BaseAnnotation) PreCheck() error {
	// rect
	if IsZeroEpsilon(b.rect.Left) && IsZeroEpsilon(b.rect.Bottom) &&
		IsZeroEpsilon(b.rect.Right) && IsZeroEpsilon(b.rect.Top) {
		return errors.New("rect must be set")
	}
	return nil
}

func (b *BaseAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// pre base check
	err := b.PreCheck()
	if err != nil {
		return err
	}

	// create annot
	annotRes, err := instance.FPDFPage_CreateAnnot(&requests.FPDFPage_CreateAnnot{
		Page:    page,
		Subtype: b.subtype,
	})
	if err != nil {
		log.Fatalf("create annot failed: %v", err)
		return err
	}
	b.annot = annotRes.Annotation

	// set title
	if b.title != "" {
		_, err = instance.FPDFAnnot_SetStringValue(&requests.FPDFAnnot_SetStringValue{
			Annotation: b.annot,
			Key:        "T",
			Value:      b.title,
		})
		if err != nil {
			log.Fatalf("set annot title failed: %v", err)
			return err
		}
	}

	// set rect
	_, err = instance.FPDFAnnot_SetRect(&requests.FPDFAnnot_SetRect{
		Annotation: b.annot,
		Rect: structs.FPDF_FS_RECTF{
			Left:   b.rect.Left,
			Bottom: b.rect.Bottom,
			Right:  b.rect.Right,
			Top:    b.rect.Top,
		},
	})
	if err != nil {
		log.Fatalf("set annot rect failed: %v", err)
		return err
	}

	// set border
	if !IsZeroEpsilon(b.Width) {
		_, err = instance.FPDFAnnot_SetBorder(&requests.FPDFAnnot_SetBorder{
			Annotation:       b.annot,
			HorizontalRadius: 0,
			VerticalRadius:   0,
			BorderWidth:      float32(b.Width),
		})
		if err != nil {
			log.Fatalf("set annot border failed: %v", err)
			return err
		}
	}

	// set strike color
	if b.strikeColor != nil {
		_, err = instance.FPDFAnnot_SetColor(&requests.FPDFAnnot_SetColor{
			Annotation: b.annot,
			ColorType:  enums.FPDFANNOT_COLORTYPE_Color,
			R:          uint(b.strikeColor.R),
			G:          uint(b.strikeColor.G),
			B:          uint(b.strikeColor.B),
			A:          uint(b.opacity),
		})
		if err != nil {
			log.Fatalf("set annot strike color failed: %v", err)
			return err
		}
	}

	// set fill color
	if b.fillColor != nil {
		_, err = instance.FPDFAnnot_SetColor(&requests.FPDFAnnot_SetColor{
			Annotation: b.annot,
			ColorType:  enums.FPDFANNOT_COLORTYPE_InteriorColor,
			R:          uint(b.fillColor.R),
			G:          uint(b.fillColor.G),
			B:          uint(b.fillColor.B),
			A:          uint(b.opacity),
		})
		if err != nil {
			log.Fatalf("set annot fill color failed: %v", err)
			return err
		}
	}	

	// set nm
	if b.nm != "" {
		_, err = instance.FPDFAnnot_SetStringValue(&requests.FPDFAnnot_SetStringValue{
			Annotation: b.annot,
			Key:        "NM",
			Value:      b.nm,
		})
		if err != nil {
			log.Fatalf("set annot nm failed: %v", err)
			return err
		}	
	}


	// set ap
	if b.ap != "" {
		log.Printf("\n\nsubtype:%s annot ap: %s\n", b.GetSubtypeName(), b.ap)

		_, err = instance.FPDFAnnot_SetAP(&requests.FPDFAnnot_SetAP{
			Annotation:     b.annot,
			AppearanceMode: enums.FPDF_ANNOT_APPEARANCEMODE_NORMAL,
			Value:          &b.ap,
		})
		if err != nil {
			log.Fatalf("set annot ap failed: %v", err)
			return err
		}
	}

	return nil
}

func (b *BaseAnnotation) GetSubtypeName() string {
	switch b.subtype {
	case enums.FPDF_ANNOT_SUBTYPE_CIRCLE:
		return "Circle"
	case enums.FPDF_ANNOT_SUBTYPE_SQUARE:
		return "Square"
	case enums.FPDF_ANNOT_SUBTYPE_CARET:
		return "Caret"
	case enums.FPDF_ANNOT_SUBTYPE_LINE:
		return "Line"
	case enums.FPDF_ANNOT_SUBTYPE_POLYGON:
		return "Polygon"
	case enums.FPDF_ANNOT_SUBTYPE_TEXT:
		return "Text"
	case enums.FPDF_ANNOT_SUBTYPE_FILEATTACHMENT:
		return "FileAttachment"
	case enums.FPDF_ANNOT_SUBTYPE_LINK:
		return "Link"
	case enums.FPDF_ANNOT_SUBTYPE_FREETEXT:
		return "FreeText"
	case enums.FPDF_ANNOT_SUBTYPE_INK:
		return "Ink"
	case enums.FPDF_ANNOT_SUBTYPE_HIGHLIGHT:
		return "Highlight"
	case enums.FPDF_ANNOT_SUBTYPE_MOVIE:
		return "Movie"
	case enums.FPDF_ANNOT_SUBTYPE_SOUND:
		return "Sound"
	case enums.FPDF_ANNOT_SUBTYPE_POPUP:
		return "Popup"
	case enums.FPDF_ANNOT_SUBTYPE_PRINTERMARK:
		return "PrinterMark"
	case enums.FPDF_ANNOT_SUBTYPE_POLYLINE:
		return "Polyline"
	case enums.FPDF_ANNOT_SUBTYPE_REDACT:
		return "Redact"
	case enums.FPDF_ANNOT_SUBTYPE_RICHMEDIA:
		return "RichMedia"
	case enums.FPDF_ANNOT_SUBTYPE_SCREEN:
		return "Screen"
	case enums.FPDF_ANNOT_SUBTYPE_WATERMARK:
		return "Watermark"
	case enums.FPDF_ANNOT_SUBTYPE_SQUIGGLY:
		return "Squiggly"
	case enums.FPDF_ANNOT_SUBTYPE_STAMP:
		return "Stamp"
	case enums.FPDF_ANNOT_SUBTYPE_STRIKEOUT:
		return "Strikeout"
	case enums.FPDF_ANNOT_SUBTYPE_THREED:
		return "ThreeD"
	case enums.FPDF_ANNOT_SUBTYPE_UNDERLINE:
		return "Underline"
	case enums.FPDF_ANNOT_SUBTYPE_TRAPNET:
		return "TrapNet"
	case enums.FPDF_ANNOT_SUBTYPE_WIDGET:
		return "Widget"
	case enums.FPDF_ANNOT_SUBTYPE_XFAWIDGET:
		return "XFAWidget"

	default:
		return "Unknown"
	}
}
