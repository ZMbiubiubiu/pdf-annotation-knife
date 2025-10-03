package annotation

import (
	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/structs"
)

func CreatePathObject(instance pdfium.Pdfium, page requests.Page, pathObjectParam *PathObjectParam) (references.FPDF_PAGEOBJECT, error) {
	// create path obj
	objRef, err := instance.FPDFPageObj_CreateNewPath(&requests.FPDFPageObj_CreateNewPath{
		X: 0,
		Y: 0,
	})
	if err != nil {
		return "", err
	}

	// stroke color
	_, err = instance.FPDFPageObj_SetStrokeColor(&requests.FPDFPageObj_SetStrokeColor{
		PageObject: objRef.PageObject,
		StrokeColor: structs.FPDF_COLOR{
			R: uint(pathObjectParam.StrikeColor.R),
			G: uint(pathObjectParam.StrikeColor.G),
			B: uint(pathObjectParam.StrikeColor.B),
			A: uint(pathObjectParam.StrikeAlpha),
		},
	})
	if err != nil {
		return "", err
	}

	// set width
	_, err = instance.FPDFPageObj_SetStrokeWidth(&requests.FPDFPageObj_SetStrokeWidth{
		PageObject:  objRef.PageObject,
		StrokeWidth: pathObjectParam.Width,
	})
	if err != nil {
		return "", err
	}

	// set line cap
	_, err = instance.FPDFPageObj_SetLineCap(&requests.FPDFPageObj_SetLineCap{
		PageObject: objRef.PageObject,
		LineCap:    pathObjectParam.LineCap,
	})
	if err != nil {
		return "", err
	}

	// set line join
	_, err = instance.FPDFPageObj_SetLineJoin(&requests.FPDFPageObj_SetLineJoin{
		PageObject: objRef.PageObject,
		LineJoin:   pathObjectParam.LineJoin,
	})
	if err != nil {
		return "", err
	}

	// join line
	for _, points := range pathObjectParam.Points {
		if len(points) == 0 {
			continue
		}
		// move to first point
		_, err = instance.FPDFPath_MoveTo(&requests.FPDFPath_MoveTo{
			PageObject: objRef.PageObject,
			X:          points[0].X,
			Y:          points[0].Y,
		})
		if err != nil {
			return "", err
		}

		for j := 1; j < len(points); j++ {
			_, err = instance.FPDFPath_LineTo(&requests.FPDFPath_LineTo{
				PageObject: objRef.PageObject,
				X:          points[j].X,
				Y:          points[j].Y,
			})
			if err != nil {
				return "", err
			}
		}
	}

	// stroke mode
	_, err = instance.FPDFPath_SetDrawMode(&requests.FPDFPath_SetDrawMode{
		PageObject: objRef.PageObject,
		FillMode:   enums.FPDF_FILLMODE_NONE,
		Stroke:     true,
	})
	if err != nil {
		return "", err
	}

	return objRef.PageObject, nil
}
