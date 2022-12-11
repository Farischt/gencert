package pdf

import (
	"fmt"
	"gencert/cert"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &PdfSaver{
		OutputDir: outputDir,
	}, nil

}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Create a background
	setPdfBackground(pdf, "assets/background.png")

	// Create a header
	setPdfHeader(pdf, "assets/gopher.png", cert.LabelCompletion)
	pdf.Ln(30)

	// Create body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(15)

	// Student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Participation
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
	pdf.Ln(15)

	setPdfFooter(pdf, "assets/stamp.png")

	// Save the pdf
	filename := fmt.Sprintf("*%s.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}

	fmt.Printf("Saved pdf certificate in %s\n", path)
	return nil
}

func setPdfBackground(pdf *gofpdf.Fpdf, imgPath string) {
	options := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()

	pdf.ImageOptions(imgPath, 0, 0, pageWidth, pageHeight, false, options, 0, "")
}

func setPdfHeader(pdf *gofpdf.Fpdf, logoPath string, header string) {
	options := gofpdf.ImageOptions{
		ImageType: "png",
	}

	// Left logo
	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	pdf.ImageOptions(logoPath, x+margin, 20, imageWidth, 0, false, options, 0, "")

	// Right logo
	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(logoPath, x-margin, 20, imageWidth, 0, false, options, 0, "")

	// Text
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, header, "C")
}

func setPdfFooter(pdf *gofpdf.Fpdf, logoPath string) {
	options := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	imageWidth := 30.0
	imageHeight := imageWidth
	pageWidth, pageHeight := pdf.GetPageSize()
	//
	x := (pageWidth / 2) + (imageWidth / 2)
	y := pageHeight - imageHeight
	pdf.ImageOptions(logoPath, x-margin, y-margin, imageWidth, 0, false, options, 0, "")

}
