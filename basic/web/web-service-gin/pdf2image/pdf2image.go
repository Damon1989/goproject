package pdf2image

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func PDFToImage(pdfPath string, outputDir string, dpi float64, scale float64) error {
	// 检查pdf文件是否存在
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		return fmt.Errorf("pdf file not exist :%s", pdfPath)
	}
	// 创建输出目录
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("create output dir error:%s", err)
	}
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return fmt.Errorf("open pdf file error:%s", err)
	}
	defer doc.Close()

	pageCount := doc.NumPage()
	fmt.Println("pageCount:", pageCount)
	for i := 0; i < pageCount; i++ {
		img, err := doc.ImageDPI(i, dpi)
		if err != nil {
			return fmt.Errorf("get page image error:%s", err)
		}
		//if scale != 1.0 {
		//	newWidth := int(float64(img.Bounds().Dx()) * scale)
		//	newHeight := int(float64(img.Bounds().Dy()) * scale)
		//	img = resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)
		//}
		pdfName := filepath.Base(pdfPath)
		pdfName = pdfName[:len(pdfName)-len(filepath.Ext(pdfName))] // 去掉.pdf后缀
		imgPath := filepath.Join(outputDir, fmt.Sprintf("%s_%d.png", pdfName, i+1))

		f, err := os.Create(imgPath)
		if err != nil {
			return fmt.Errorf("创建图片文件失败: %v", err)
		}
		defer f.Close()

		// 保存为PNG格式（也可保存为JPG，需引入image/jpeg库）
		if err := png.Encode(f, img); err != nil {
			return fmt.Errorf("保存第%d页图片失败: %v", i+1, err)
		}

		fmt.Printf("第%d页转换完成：%s\n", i+1, imgPath)
	}
	fmt.Println("PDF转图片完成！")
	return nil
}
