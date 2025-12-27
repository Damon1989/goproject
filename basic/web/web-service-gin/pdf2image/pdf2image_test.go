package pdf2image

import "testing"

func TestPDFToImage(t *testing.T) {
	PDFToImage("D:\\tmp\\pdf\\1.pdf", "D:\\tmp\\pdf\\1", 300, 1)
}
