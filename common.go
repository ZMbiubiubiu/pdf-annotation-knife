package annotation

import (
	"context"
	"image/color"
	"log"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/structs"
)

// annotation width
type Width float32

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

type BaseAnnotation struct {
	Page        requests.Page
	Annotation  references.FPDF_ANNOTATION
	Subtype     enums.FPDF_ANNOTATION_SUBTYPE
	Rect        Rect
	Width       Width
	StrikeColor *color.RGBA
	FillColor   *color.RGBA
	AP          string
}

func (b *BaseAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {

	// create annot
	annotRes, err := instance.FPDFPage_CreateAnnot(&requests.FPDFPage_CreateAnnot{
		Page:    b.Page,
		Subtype: b.Subtype,
	})
	if err != nil {
		log.Fatalf("create annot failed: %v", err)
		return err
	}
	b.Annotation = annotRes.Annotation

	// set rect
	_, err = instance.FPDFAnnot_SetRect(&requests.FPDFAnnot_SetRect{
		Annotation: b.Annotation,
		Rect: structs.FPDF_FS_RECTF{
			Left:   b.Rect.Left,
			Bottom: b.Rect.Bottom,
			Right:  b.Rect.Right,
			Top:    b.Rect.Top,
		},
	})
	if err != nil {
		log.Fatalf("set annot rect failed: %v", err)
		return err
	}

	// set border
	_, err = instance.FPDFAnnot_SetBorder(&requests.FPDFAnnot_SetBorder{
		Annotation:       b.Annotation,
		HorizontalRadius: 0,
		VerticalRadius:   0,
		BorderWidth:      float32(b.Width),
	})
	if err != nil {
		log.Fatalf("set annot border failed: %v", err)
		return err
	}

	// set strike color
	if b.StrikeColor != nil {
		_, err = instance.FPDFAnnot_SetColor(&requests.FPDFAnnot_SetColor{
			Annotation: b.Annotation,
			ColorType:  enums.FPDFANNOT_COLORTYPE_Color,
			R:          uint(b.StrikeColor.R),
			G:          uint(b.StrikeColor.G),
			B:          uint(b.StrikeColor.B),
			A:          uint(b.StrikeColor.A),
		})
		if err != nil {
			log.Fatalf("set annot strike color failed: %v", err)
			return err
		}
	}

	// set fill color
	if b.FillColor != nil {
		_, err = instance.FPDFAnnot_SetColor(&requests.FPDFAnnot_SetColor{
			Annotation: b.Annotation,
			ColorType:  enums.FPDFANNOT_COLORTYPE_InteriorColor,
			R:          uint(b.FillColor.R),
			G:          uint(b.FillColor.G),
			B:          uint(b.FillColor.B),
			A:          uint(b.FillColor.A),
		})
		if err != nil {
			log.Fatalf("set annot fill color failed: %v", err)
			return err
		}
	}

	// set ap
	if b.AP != "" {
		_, err = instance.FPDFAnnot_SetAP(&requests.FPDFAnnot_SetAP{
			Annotation:     b.Annotation,
			AppearanceMode: enums.FPDF_ANNOT_APPEARANCEMODE_NORMAL,
			Value:          &b.AP,
		})
		if err != nil {
			log.Fatalf("set annot ap failed: %v", err)
			return err
		}
	}

	return nil
}
