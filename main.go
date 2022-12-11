package main

import (
	"flag"
	"fmt"
	"gencert/cert"
	"gencert/html"

	"gencert/pdf"
	"os"
)

const outputDir string = "output"

func main() {
	fileExtension := flag.String("type", "pdf", "the extension of the certificate")
	csvFile := flag.String("filename", "", "csv file input")
	flag.Parse()

	if len(*csvFile) == 0 {
		fmt.Printf("Csv file %s is empty !\n", *csvFile)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error

	switch *fileExtension {
	case "pdf":
		saver, err = pdf.New(fmt.Sprintf("%s/pdf", outputDir))
		if err != nil {
			handleErrorAndExit(err)
		}

	case "html":
		saver, err = html.New(fmt.Sprintf("%s/html", outputDir))
		if err != nil {
			handleErrorAndExit(err)
		}

	default:
		fmt.Printf("Invalid extension type %s. We currently provide html & pdf certificates only !\n", *fileExtension)
		os.Exit(1)
	}

	certs, err := cert.ParseCsvFile(*csvFile)
	if err != nil {
		handleErrorAndExit(err)
	}

	for _, c := range certs {
		err := saver.Save(*c)
		if err != nil {
			fmt.Printf("Couldn't save certificate for student %s\n", c.Name)
			fmt.Printf("%s\n", err.Error())
		}
	}

}

func handleErrorAndExit(err error) {
	fmt.Printf("%s\n", err.Error())
	os.Exit(1)
}
