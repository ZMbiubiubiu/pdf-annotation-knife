package annotation

import (
	"errors"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/structs"
)

// GetImageDimensions 从指定路径获取图片的宽度和高度
func GetImageDimensions(imagePath string) (width int, height int, err error) {
	// 1. 打开文件
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close() // 确保文件关闭

	// 2. 解码图片
	// image.Decode 函数会自动识别格式并返回 image.Image 接口
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, fmt.Errorf("无法解码图片: %w", err)
	}

	// 3. 获取图片的边界信息
	// Bounds() 方法返回一个 image.Rectangle 结构体
	bounds := img.Bounds()

	// 宽度 = X 轴的最大值 - 最小值 (Max.X - Min.X)
	width = bounds.Dx()
	// 或者直接使用 width = bounds.Max.X - bounds.Min.X

	// 高度 = Y 轴的最大值 - 最小值 (Max.Y - Min.Y)
	height = bounds.Dy()
	// 或者直接使用 height = bounds.Max.Y - bounds.Min.Y

	log.Printf("image dimensions: %d x %d", width, height)

	return width, height, nil
}

// CreateImgObject creates an image object.
func CreateImgObject(instance pdfium.Pdfium, rect Rect, imgParam *ImageObjectParam) (references.FPDF_PAGEOBJECT, error) {
	var imgRef references.FPDF_PAGEOBJECT
	var err error

	switch imgParam.ImgType {
	case "jpg", "jpeg":
		imgRef, err = createJPEGImgObject(instance, imgParam.Document, imgParam.FilePath)
	case "png":
		// TODO: create png image object
	default:
		return "", errors.New("unsupported image type")
	}

	if err != nil {
		return "", err
	}

	// set matrix
	_, err = instance.FPDFImageObj_SetMatrix(&requests.FPDFImageObj_SetMatrix{
		ImageObject: imgRef,
		Transform: structs.FPDF_FS_MATRIX{
			A: rect.Right - rect.Left,
			B: 0,
			C: 0,
			D: rect.Top - rect.Bottom,
			E: rect.Left,
			F: rect.Bottom,
		},
	})
	if err != nil {
		return "", err
	}

	return imgRef, nil
}

// createJPEGImgObject creates a jpeg image object.
func createJPEGImgObject(instance pdfium.Pdfium, doc references.FPDF_DOCUMENT, filePath string) (references.FPDF_PAGEOBJECT, error) {
	// create image object
	imgRef, err := instance.FPDFPageObj_NewImageObj(&requests.FPDFPageObj_NewImageObj{
		Document: doc,
	})
	if err != nil {
		return "", err
	}

	// load jpeg file
	_, err = instance.FPDFImageObj_LoadJpegFile(&requests.FPDFImageObj_LoadJpegFile{
		ImageObject: imgRef.PageObject,
		FilePath:    filePath,
	})
	if err != nil {
		return "", err
	}

	return imgRef.PageObject, nil
}
