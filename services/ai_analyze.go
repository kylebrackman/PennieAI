package services

import (
	"PennieAI/utils"
	"mime/multipart"
)

func AnalyzeDocument(file *multipart.FileHeader) ([]utils.Window, error) {
	// Get file lines
	fileLines, err := utils.GetFileLines(file)
	if err != nil {
		return nil, err
	}

	// Now segment the document into windows

	windows := utils.WindowBuilder(fileLines, nil)
	// ... continue processing

	return windows, nil

}
