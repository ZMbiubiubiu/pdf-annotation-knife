// 自由文本
package annotation

import (
	"context"
	"fmt"
	"image/color"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

var (
	DefaultFontSize  = 12
	DefaultFontColor = color.RGBA{0, 0, 0, 255}
)

type FreeTextAnnotation struct {
	BaseAnnotation
	Contents  string
	FontColor *color.RGBA
	FontSize  int
}

func NewFreeTextAnnotation(page requests.Page) *FreeTextAnnotation {
	return &FreeTextAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page:    page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_FREETEXT,
		},
	}
}

func (f *FreeTextAnnotation) GenerateAppearance() error {
	// todo generate freetext appearance
	f.AP = ""
	return nil
}

func (f *FreeTextAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	// set default font size and color
	if f.FontSize == 0 {
		f.FontSize = DefaultFontSize
	}
	if f.FontColor == nil {
		f.FontColor = &DefaultFontColor
	}

	// create annotation
	err := f.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}

	// set contents
	_, err = instance.FPDFAnnot_SetStringValue(&requests.FPDFAnnot_SetStringValue{
		Annotation: f.Annotation,
		Key:        "Contents",
		Value:      f.Contents,
	})
	if err != nil {
		return err
	}

	// set font color
	da := fmt.Sprintf("%d Tf %.3f %.3f %.3f rg ", 12, float32(f.FontColor.R)/255, float32(f.FontColor.G)/255, float32(f.FontColor.B)/255)
	_, err = instance.FPDFAnnot_SetStringValue(&requests.FPDFAnnot_SetStringValue{
		Annotation: f.Annotation,
		Key:        "DA",
		Value:      da,
	})
	if err != nil {
		return err
	}
	return nil
}
