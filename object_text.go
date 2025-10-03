package annotation

import (
	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
)

func CreateTextObject(instance pdfium.Pdfium, textParam *TextObjectParam) (references.FPDF_PAGEOBJECT, error) {

	// create text object
	textRef, err := instance.FPDFPageObj_CreateTextObj(&requests.FPDFPageObj_CreateTextObj{
		// Document: nil,
		// Font:     nil,
		FontSize: textParam.FontSize,
	})
	if err != nil {
		return "", err
	}

	// set text

	return textRef.PageObject, nil
}
