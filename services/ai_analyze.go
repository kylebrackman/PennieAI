package services

import (
	"PennieAI/utils"
	"mime/multipart"
)

func AnalyzeDocument(file *multipart.FileHeader) ([]string, error) {
	// Get file lines
	fileLines, err := utils.GetFileLines(file)
	if err != nil {
		return nil, err
	}

	return fileLines, nil

	// ... continue processing
}
