package annotation

import (
	"context"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/structs"
)

type Point struct {
	X,Y float32
}

type InkAnnotation struct {
	BaseAnnotation
	Points [][]Point
}

func NewInkAnnotation(page requests.Page) *InkAnnotation {
	return &InkAnnotation{
		BaseAnnotation: BaseAnnotation{
			Page: page,
			Subtype: enums.FPDF_ANNOT_SUBTYPE_INK,
		},
	}
}

func (i *InkAnnotation) GenerateAppearance() error {
	// todo generate ink appearance
	i.AP = ""
	return nil
}

func (i *InkAnnotation) convertPointToPdfiumFormat(points []Point) []structs.FPDF_FS_POINTF {
	pdfiumPoints := make([]structs.FPDF_FS_POINTF, 0, len(points))
	for _, point := range points {
		pdfiumPoints = append(pdfiumPoints, structs.FPDF_FS_POINTF{
			X: point.X,
			Y: point.Y,
		})
	}
	return pdfiumPoints
}

func (i *InkAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium) error {
	err := i.BaseAnnotation.AddAnnotationToPage(ctx, instance)
	if err != nil {
		return err
	}
	// 插入点
	for _, points := range i.Points {
		p := i.convertPointToPdfiumFormat(points)
		_, err = instance.FPDFAnnot_AddInkStroke(&requests.FPDFAnnot_AddInkStroke{
			Annotation: i.Annotation,
			Points:     p,
		})
		if err != nil {
			return err
		}
	}
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: i.Annotation,
	})
	if err != nil {
		return err
	}
	return nil
}