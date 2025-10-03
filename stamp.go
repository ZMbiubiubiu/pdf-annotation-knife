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

// PathObjectParam is the parameter for creating a path object.
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

// ImageObjectParam is the parameter for creating an image object.
type ImageObjectParam struct {
	ImgType  string // png/jpg/jpeg
	Document references.FPDF_DOCUMENT
	FilePath string
}

func (s *StampAnnotation) SetImgObject(imgType string, document references.FPDF_DOCUMENT, filePath string) {
	s.objectType = StampObjectImg
	s.imgObject = &ImageObjectParam{
		ImgType:  imgType,
		Document: document,
		FilePath: filePath,
	}
}

// TextObjectParam is the parameter for creating a text object.
type TextObjectParam struct {
	Document references.FPDF_DOCUMENT
	FontSize float32
	Text     string
}

type StampAnnotation struct {
	BaseAnnotation
	objectType StampObjectType
	pathObject *PathObjectParam
	imgObject  *ImageObjectParam
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
		// TODO objRef, err = CreateTextObject(instance, page, s.textObject)
	case StampObjectImg:
		objRef, err = CreateImgObject(instance, s.rect, s.imgObject)
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
