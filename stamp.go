package annotation

import (
	"context"
	"errors"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
)

type StampObjectType int

const (
	StampObjectPath StampObjectType = 1
	StampObjectText StampObjectType = 2
	StampObjectImg  StampObjectType = 3
)

type PathObjectParam struct {
	Points      [][]Point
	Width       float32
	StrikeColor Color
	StrikeAlpha uint8
	LineCap     enums.FPDF_LINECAP
	LineJoin    enums.FPDF_LINEJOIN
}

func (s *StampAnnotation) SetPathObject(points [][]Point, width float32, strikeColor Color, strikeAlpha uint8) {
	s.objectType = StampObjectPath
	s.pathObject = &PathObjectParam{
		Points:      points,
		Width:       width,
		StrikeColor: strikeColor,
		StrikeAlpha: strikeAlpha,
		LineCap:     enums.FPDF_LINECAP_ROUND,
		LineJoin:    enums.FPDF_LINEJOIN_ROUND,
	}
}

type StampAnnotation struct {
	BaseAnnotation
	objectType StampObjectType
	pathObject *PathObjectParam
}

func NewStampAnnotation() *StampAnnotation {
	return &StampAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_STAMP,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
	}
}

func (s *StampAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	// create annotation
	err := s.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// create page object
	var objRef references.FPDF_PAGEOBJECT
	switch s.objectType {
	case StampObjectPath:
		objRef, err = CreatePathObject(instance, page, s.pathObject)
	case StampObjectText:
		// objRef, err = CreateTextObject(instance, page, s.textObject)
	case StampObjectImg:
		// objRef, err = CreateImgObject(instance, page, s.imgObject)
	default:
		return errors.New("object type not supported")
	}

	if err != nil {
		return err
	}

	// insert object
	_, err = instance.FPDFAnnot_AppendObject(&requests.FPDFAnnot_AppendObject{
		Annotation: s.annot,
		PageObject: objRef,
	})
	if err != nil {
		return err
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: s.annot,
	})
	if err != nil {
		return err
	}

	return nil
}
