package services

import (
	"PennieAI/models"
	"PennieAI/prompts"
	"PennieAI/utils"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"
)

func AnalyzeDocument(file *multipart.FileHeader, aiService *AIService) (*models.Patient, []models.AnalyzedDocument, error) {

	var patient models.Patient
	var analyzedDocuments []models.AnalyzedDocument

	fileLines, err := utils.GetFileLines(file)
	if err != nil {
		return nil, nil, err
	}

	windows := utils.WindowBuilder(fileLines, nil)

	for _, window := range windows {
		// Build incremental notice if we have previous documents
		// Build incremental notice if we have previous documents
		// This tells OpenAI what we've already found in earlier windows to avoid duplicates
		var incrementalNotice string

		var promptBuilder strings.Builder
		promptBuilder.WriteString(prompts.BasePrompt)

		if len(analyzedDocuments) > 0 {
			patientJSON, _ := json.MarshalIndent(patient, "  ", "  ")
			docsJSON, _ := json.MarshalIndent(analyzedDocuments, "  ", "  ")
			incrementalNotice = fmt.Sprintf(prompts.IncrementalNoticeTemplate, patientJSON, docsJSON)

			promptBuilder.WriteString("\n")

			promptBuilder.WriteString(incrementalNotice)

			promptBuilder.WriteString("\n")
		}
		promptBuilder.WriteString("Here is the text chunk:\n")

		for lineIndex, line := range window.WindowLines {
			lineNumber := window.StartIndex + lineIndex + 1
			promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", lineNumber, line))
		}

		analyzedDocuments = append(analyzedDocuments, models.AnalyzedDocument{
			Title: fmt.Sprintf("test_%d", window.StartIndex),
			// Just for testing purposes
			StartLine: 0,
			EndLine:   1,
		})

		response, err := aiService.Query(context.Background(), promptBuilder.String(), nil)

		if err != nil {
			return nil, nil, fmt.Errorf("AI query failed: %w", err)
		}

		fmt.Printf("OpenAI Response: %+v\n", response)

		// See Q&A 2025-10-14 for more info on this syntax
		if patientData, ok := response["patient"].(map[string]interface{}); ok {
			// Merge patient fields (prefer non-empty values)
			if name, ok := patientData["name"].(string); ok && name != "" {
				patient.Name = name
			}
			if species, ok := patientData["possibleSpecies"].(string); ok && species != "" {
				// First, check if the pointer is nil
				if patient.PossibleSpecies == nil {
					// If it's nil, initialize it with a pointer to an empty slice
					patient.PossibleSpecies = &[]string{}
				}

				newPossibleSpecies := true

				for _, possibleSpecies := range *patient.PossibleSpecies {
					if possibleSpecies == species {
						newPossibleSpecies = false
					}
				}
				if newPossibleSpecies {
					*patient.PossibleSpecies = append(*patient.PossibleSpecies, species)
				}
			}
			if breed, ok := patientData["possibleBreed"].(string); ok && breed != "" {
				if patient.PossibleBreed == nil {
					patient.PossibleBreed = &[]string{}
				}

				newPossibleBreed := true

				for _, possibleBreed := range *patient.PossibleBreed {
					if possibleBreed == breed {
						newPossibleBreed = false
					}
				}

				if newPossibleBreed {
					*patient.PossibleBreed = append(*patient.PossibleBreed, breed)
				}
			}
			if sex, ok := patientData["sex"].(string); ok && sex != "" {
				patient.Sex = &sex
			}
		}

		if docs, ok := response["documents"].([]interface{}); ok {
			for _, doc := range docs {
				if docDetails, ok := doc.(map[string]interface{}); ok {
					// Check if this document already exists (deduplicate by start_line)
					startLine := int64(docDetails["start_line"].(float64))
					title := docDetails["title"].(string)
					endLine := int64(docDetails["end_line"].(float64))
					numberOfLines := endLine - startLine

					isDuplicate := false
					for _, existingDoc := range analyzedDocuments {
						if existingDoc.StartLine == startLine {
							isDuplicate = true
							break
						}
					}

					windowStartLine := startLine - int64(window.StartIndex)
					windowEndLine := endLine - int64(window.StartIndex)

					if !isDuplicate {
						analyzedDocuments = append(analyzedDocuments, models.AnalyzedDocument{
							Title:         title,
							StartLine:     startLine,
							EndLine:       endLine,
							NumberOfLines: numberOfLines,
							WindowLines:   window.WindowLines[windowStartLine:windowEndLine],
						})
					}
				}
			}
		}
	}

	return &patient, analyzedDocuments, nil
}
