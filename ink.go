package annotation

import (
	"context"
	"fmt"
	"strings"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/requests"
)

type InkAnnotation struct {
	BaseAnnotation
	LineStyle
	Points [][]Point
}

func NewInkAnnotation() *InkAnnotation {
	return &InkAnnotation{
		BaseAnnotation: BaseAnnotation{
			subtype: enums.FPDF_ANNOT_SUBTYPE_INK,
			nm:      GenerateUUID(),
			opacity: DefaultOpacity,
		},
		LineStyle: LineStyle{
			StrikeLineCap:  enums.FPDF_LINECAP_ROUND,
			StrikeLineJoin: enums.FPDF_LINEJOIN_ROUND,
		},
	}
}

func (i *InkAnnotation) GenerateAppearance() error {
	// generate ink appearance
	i.ap = strings.Join([]string{
		i.GetColorAP(),
		i.GetWidthAP(),
		i.GetPDFOpacityAP(),
		i.GetLineStyleAP(),
		i.pointsCallback(),
	}, "\n")

	return nil
}

func (i *InkAnnotation) pointsCallback() string {
	var path string
	for _, points := range i.Points {
		for i, point := range points {
			op := "l"
			if i == 0 {
				op = "m"
			}
			path += fmt.Sprintf("%.3f %.3f %s ", point.X, point.Y, op)
		}
		path += "S "
	}
	return path
}

func (i *InkAnnotation) AddAnnotationToPage(ctx context.Context, instance pdfium.Pdfium, page requests.Page) error {
	err := i.BaseAnnotation.AddAnnotationToPage(ctx, instance, page)
	if err != nil {
		return err
	}

	// insert points
	for _, points := range i.Points {
		p := convertPointToPdfiumFormat(points)
		_, err = instance.FPDFAnnot_AddInkStroke(&requests.FPDFAnnot_AddInkStroke{
			Annotation: i.annot,
			Points:     p,
		})
		if err != nil {
			return err
		}
	}

	// close annotation
	_, err = instance.FPDFPage_CloseAnnot(&requests.FPDFPage_CloseAnnot{
		Annotation: i.annot,
	})
	if err != nil {
		return err
	}
	return nil
}
