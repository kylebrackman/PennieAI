package utils

import (
	"io"
	"mime/multipart"
	"strings"
)

func GetFileLines(file *multipart.FileHeader) ([]string, error) {
	openedFile, err := file.Open() // This returns an io.ReadCloser, which is a reader interface
	// See Q&A 2025-10-11 for more info
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()

	// Read the content into bytes
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		return nil, err
	}

	// Convert bytes to string
	fileContent := string(fileBytes)

	// Now you can split it
	return strings.Split(fileContent, "\n"), nil
}
