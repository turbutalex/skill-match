package parser

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"os"
)

type ParsedData struct {
	Name   string
	Email  string
	Phone  string
	Skills []string
}

func Extract(path string) (string, ParsedData, error) {

	f, err := os.Open(path)
	if err != nil {
		return "", ParsedData{}, err
	}

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return "", ParsedData{}, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", ParsedData{}, err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")

	var resumeRawText string

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", ParsedData{}, err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", ParsedData{}, err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return "", ParsedData{}, err
		}

		resumeRawText += text
	}

	resume, err := SendRequest(resumeRawText)
	if err != nil {
		return "", ParsedData{}, err
	}

	parsed := ParsedData{
		Name:   resume.ParsedName,
		Email:  resume.Email,
		Phone:  resume.Phone,
		Skills: resume.Skills,
	}

	return resumeRawText, parsed, nil
}
