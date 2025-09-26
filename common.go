package annotation

import (
	"context"
	"image/color"

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

type BaseAnnotation struct {
	Annotation  references.FPDF_ANNOTATION
	Subtype     enums.FPDF_ANNOTATION_SUBTYPE
	Rect        Rect
	Width       Width
	StrikeColor *color.RGBA
	FillColor   *color.RGBA
	AP          string
}

func (b *BaseAnnotation) CreateAnnotation(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {

	// create annot
	annotRes, err := instance.FPDFPage_CreateAnnot(&requests.FPDFPage_CreateAnnot{
		Page:    page,
		Subtype: b.Subtype,
	})
	if err != nil {
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
			return err
		}
	}

	return nil
}
