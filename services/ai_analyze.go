package services

import (
	"PennieAI/utils"
	"mime/multipart"
	"strconv"
)

func AnalyzeDocument(file *multipart.FileHeader) ([]utils.Window, error) {
	// Get file lines
	fileLines, err := utils.GetFileLines(file)
	if err != nil {
		return nil, err
	}

	windows := utils.WindowBuilder(fileLines, nil)

	for _, window := range windows {
		for lineIndex, line := range window.WindowLines {
			lineNumber := window.StartIndex + lineIndex + 1 // +1 for 1-based line numbers
			window.WindowLines[lineIndex] = strconv.Itoa(lineNumber) + ": " + line
		}
	}
	return windows, nil

}
