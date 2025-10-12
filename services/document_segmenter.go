package services

import (
	"PennieAI/utils"
	"strings"
)

func DocumentSegmenter(file) {
	// Use the WindowBuilder from utils to segment the document

	fileLines := strings.Split(fileContent, "\n")

	builder := utils.NewWindowBuilder()

}
