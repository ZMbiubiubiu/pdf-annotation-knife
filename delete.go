package annotation

import (
	"errors"
	"sort"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
)

type DeleteType int

const (
	DeleteByNM    = 1 // delete by unique name
	DeleteByIndex = 2 // delete by index
	DeleteByPage  = 3 // delete all Annot in given pages
	DeleteAll     = 4 // delete all Annot(every page)
)

type DeleteOnePageAnnot struct {
	PageNumber   int      // page num, start from 0
	AnnotNMs     []string // required when DeleteType is DeleteByNM
	AnnotIndices []int    // required when DeleteType is DeleteByIndex
}

type DeleteAnnot struct {
	DeleteType         DeleteType
	DeleteOnePageAnnot []DeleteOnePageAnnot
}

// DeleteAnnotInPDF delete Annot in a pdf
// ps: need load pdf first
func DeleteAnnotInPDF(instance pdfium.Pdfium, pdfFilePath, password string, deleteAnnot DeleteAnnot) (int, error) {
	// load pdf
	pdfDoc, err := instance.FPDF_LoadDocument(&requests.FPDF_LoadDocument{
		Path:     &pdfFilePath,
		Password: &password,
	})
	if err != nil {
		return 0, err
	}
	defer instance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
		Document: pdfDoc.Document,
	})
	if err != nil {
		return 0, err
	}

	return DeleteAnnotInPDFV2(instance, pdfDoc.Document, deleteAnnot)
}

// DeleteAnnotInPDFV2 delete Annot in a pdf
func DeleteAnnotInPDFV2(instance pdfium.Pdfium, pdfDoc references.FPDF_DOCUMENT, deleteAnnot DeleteAnnot) (deleted int, err error) {
	// get page count
	pageCount, err := instance.FPDF_GetPageCount(&requests.FPDF_GetPageCount{
		Document: pdfDoc,
	})
	if err != nil {
		return 0, err
	}

	if pageCount.PageCount == 0 {
		return 0, errors.New("pdf has no page")
	}

	// TODO: validate delete Annot

	// delete Annot
	switch deleteAnnot.DeleteType {
	case DeleteByIndex:
		deleted, err = DeleteAnnotByIndexs(instance, pdfDoc, deleteAnnot.DeleteOnePageAnnot)
	case DeleteByNM:
		deleted, err = DeleteAnnotByNMs(instance, pdfDoc, deleteAnnot.DeleteOnePageAnnot)
	case DeleteByPage:
		deleted, err = DeleteAllAnnotInGivenPage(instance, pdfDoc, deleteAnnot.DeleteOnePageAnnot)
	case DeleteAll:
		deleted, err = DeleteAllAnnotInEveryPage(instance, pdfDoc, pageCount.PageCount)
	default:
		err = errors.New("invalid delete type")
	}

	if err != nil {
		return 0, err
	}

	return deleted, nil
}

// DeleteAnnotByIndexs delete Annot in a pdf by indexs
func DeleteAnnotByIndexs(instance pdfium.Pdfium, pdfDoc references.FPDF_DOCUMENT, deleteAnnot []DeleteOnePageAnnot) (deleted int, err error) {
	for _, item := range deleteAnnot {
		// build page request
		page := requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: pdfDoc,
				Index:    item.PageNumber,
			},
		}
		// delete Annot in page by indexs
		num, err := deleteAnnotInPageByIndexs(instance, page, item.AnnotIndices)
		if err != nil {
			return deleted, err
		}
		deleted += num
	}
	return deleted, nil
}

// deleteAnnotInPageByIndexs delete Annot in a page by indexs
func deleteAnnotInPageByIndexs(instance pdfium.Pdfium, page requests.Page, indexs []int) (deleted int, err error) {
	// sort index, from big to small
	sort.Sort(sort.Reverse(sort.IntSlice(indexs)))

	for _, index := range indexs {
		_, err := instance.FPDFPage_RemoveAnnot(&requests.FPDFPage_RemoveAnnot{
			Page:  page,
			Index: index,
		})
		if err != nil {
			return deleted, err
		}
		deleted++
	}

	return deleted, nil
}

// GetAnnotNM get Annot names in a pdf by page numbers
// res: map[pageNum]map[AnnotNM]index
func GetAnnotNM(instance pdfium.Pdfium, pdfDoc references.FPDF_DOCUMENT, pageNums []int) (map[int]map[string]int, error) {

	res := make(map[int]map[string]int)

	for _, pageNum := range pageNums {
		page := requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: pdfDoc,
				Index:    pageNum,
			},
		}
		Annot, err := instance.FPDFPage_GetAnnotCount(&requests.FPDFPage_GetAnnotCount{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		if Annot.Count == 0 {
			res[pageNum] = make(map[string]int) // empty map
			continue
		}
		AnnotNames := make(map[string]int, Annot.Count)
		for i := 0; i < Annot.Count; i++ {
			annotRes, err := instance.FPDFPage_GetAnnot(&requests.FPDFPage_GetAnnot{
				Page:  page,
				Index: i,
			})
			if err != nil {
				return nil, err
			}

			name, err := instance.FPDFAnnot_GetStringValue(&requests.FPDFAnnot_GetStringValue{
				Annotation: annotRes.Annotation,
				Key:        "NM",
			})
			if err != nil {
				return nil, err
			}
			AnnotNames[name.Value] = i
		}
		res[pageNum] = AnnotNames
	}

	return res, nil
}

func DeleteAnnotByNMs(instance pdfium.Pdfium, pdfDoc references.FPDF_DOCUMENT, deleteAnnot []DeleteOnePageAnnot) (int, error) {

	// step1. get page nums
	var pageNums = make([]int, 0, len(deleteAnnot))
	for _, item := range deleteAnnot {
		pageNums = append(pageNums, item.PageNumber)
	}

	// step2. get Annot nm in given pages
	nms, err := GetAnnotNM(instance, pdfDoc, pageNums)
	if err != nil {
		return 0, err
	}

	// step3. convert Annot nm to index
	var deleteAnnotByIndex = make([]DeleteOnePageAnnot, len(deleteAnnot))
	for _, item := range deleteAnnot {
		var indexs = make([]int, 0, len(item.AnnotNMs))
		for _, nm := range item.AnnotNMs {
			if index, ok := nms[item.PageNumber][nm]; ok {
				indexs = append(indexs, index)
			}
		}
		deleteAnnotByIndex = append(deleteAnnotByIndex, DeleteOnePageAnnot{
			PageNumber:   item.PageNumber,
			AnnotIndices: indexs,
		})
	}

	// step4. delete Annot by index
	return DeleteAnnotByIndexs(instance, pdfDoc, deleteAnnotByIndex)
}

// DeleteAllAnnotInGivenPage delete Annot in given pages
func DeleteAllAnnotInGivenPage(instance pdfium.Pdfium, pdfDoc references.FPDF_DOCUMENT, deleteAnnot []DeleteOnePageAnnot) (int, error) {
	var totalDeleted int
	for _, item := range deleteAnnot {
		page := requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: pdfDoc,
				Index:    item.PageNumber,
			},
		}
		num, err := deleteAllAnnotInPage(instance, page)
		if err != nil {
			return 0, err
		}
		totalDeleted += num
	}
	return totalDeleted, nil
}

// deleteAllAnnotInPage delete all Annot in a page
func deleteAllAnnotInPage(instance pdfium.Pdfium, page requests.Page) (deleted int, err error) {
	annotCount, err := instance.FPDFPage_GetAnnotCount(&requests.FPDFPage_GetAnnotCount{
		Page: page,
	})
	if err != nil {
		return 0, err
	}

	for i := annotCount.Count - 1; i >= 0; i-- {
		_, err := instance.FPDFPage_RemoveAnnot(&requests.FPDFPage_RemoveAnnot{
			Page:  page,
			Index: i,
		})
		if err != nil {
			return deleted, err
		}
		deleted++
	}

	return deleted, nil
}

// DeleteAllAnnotInEveryPage delete all Annot in every page of a pdf
func DeleteAllAnnotInEveryPage(instance pdfium.Pdfium, pdfDoc references.FPDF_DOCUMENT, pageCount int) (int, error) {
	var totalDeleted int
	for i := 0; i < pageCount; i++ {
		page := requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: pdfDoc,
				Index:    i,
			},
		}
		num, err := deleteAllAnnotInPage(instance, page)
		if err != nil {
			return 0, err
		}
		totalDeleted += num
	}
	return totalDeleted, nil
}
