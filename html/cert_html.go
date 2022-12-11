package html

import (
	"fmt"
	"gencert/cert"
	"os"
	"path"
	"text/template"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputDir string) (*HtmlSaver, error) {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &HtmlSaver{
		OutputDir: outputDir,
	}, nil
}

func (p *HtmlSaver) Save(cert cert.Cert) error {
	template, err := template.New("certificate").Parse(HtmlCertificateTemplate)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("*%s.html", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = template.Execute(file, cert)
	if err != nil {
		return err
	}

	fmt.Printf("Saved html certificate in %s\n", path)
	return nil

}
