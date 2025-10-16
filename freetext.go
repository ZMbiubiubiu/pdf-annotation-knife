// 自由文本
package annotation

import (
	"context"
	"fmt"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

var (
	DefaultFontSize  = 12
	DefaultFontColor = Color{R: 0, G: 0, B: 0}
)

type FreeTextAnnotation struct {
	BaseAnnotation
	Contents  string
	FontColor Color
	FontSize  int
}

func NewFreeTextAnnotation() *FreeTextAnnotation {
	return &FreeTextAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_FREETEXT,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
		FontSize:  DefaultFontSize,
		FontColor: DefaultFontColor,
	}
}

func (f *FreeTextAnnotation) SetFontColor(color Color) {
	f.FontColor = color
}

func (f *FreeTextAnnotation) SetFontSize(size int) {
	f.FontSize = size
}

func (f *FreeTextAnnotation) GenerateAppearance() error {
	// TODO generate freetext appearance
	f.ap = ""
	return nil
}

func (f *FreeTextAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// set default font size and color
	if f.FontSize == 0 {
		f.FontSize = DefaultFontSize
	}

	// create annotation
	err := f.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// set contents
	_, err = instance.FPDFAnnot_SetStringValue(&requests.FPDFAnnot_SetStringValue{
		Annotation: f.annot,
		Key:        "Contents",
		Value:      f.Contents,
	})
	if err != nil {
		return err
	}

	// set font color
	da := fmt.Sprintf("%d Tf %.3f %.3f %.3f rg ", 12, float32(f.FontColor.R)/255, float32(f.FontColor.G)/255, float32(f.FontColor.B)/255)
	_, err = instance.FPDFAnnot_SetStringValue(&requests.FPDFAnnot_SetStringValue{
		Annotation: f.annot,
		Key:        "DA",
		Value:      da,
	})
	if err != nil {
		return err
	}
	return nil
}
