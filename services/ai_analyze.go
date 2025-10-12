package services

import (
	"io"
	"mime/multipart"
	"strings"
)

func AnalyzeDocument(file *multipart.FileHeader) error {
	// Open the file
	openedFile, err := file.Open() // This returns an io.ReadCloser, which is a reader interface
	// See Q&A 2025-10-11 for more info
	if err != nil {
		return err
	}
	defer openedFile.Close()

	// Read the content into bytes
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		return err
	}

	// Convert bytes to string
	fileContent := string(fileBytes)

	// Now you can split it
	return strings.Split(fileContent, "\n")

	// ... continue processing
}
