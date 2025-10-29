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

	// Get file lines
	fileLines, err := utils.GetFileLines(file)
	if err != nil {
		return nil, nil, err
	}

	windows := utils.WindowBuilder(fileLines, nil)

	// Process each window
	for _, window := range windows {
		// Build incremental notice if we have previous documents
		// Build incremental notice if we have previous documents
		// This tells OpenAI what we've already found in earlier windows to avoid duplicates
		var incrementalNotice string

		var promptBuilder strings.Builder
		promptBuilder.WriteString(prompts.BasePrompt)

		if len(analyzedDocuments) > 0 {

			// Convert patient struct to pretty JSON string
			// Example: {"name": "Bella", "species": "Dog", "breed": "Golden Retriever"}
			patientJSON, _ := json.MarshalIndent(patient, "  ", "  ")
			// Convert documents slice to pretty JSON array string
			// Example: [{"title": "Lab Report", "start_line": 1, "end_line": 45}, ...]
			docsJSON, _ := json.MarshalIndent(analyzedDocuments, "  ", "  ")
			incrementalNotice = fmt.Sprintf(prompts.IncrementalNoticeTemplate, patientJSON, docsJSON)

			promptBuilder.WriteString("\n")

			promptBuilder.WriteString(incrementalNotice)

			// Insert the JSON strings into the template
			// The template has two %s placeholders - one for patient, one for documents
			// This creates a message like:
			// "Here's the current patient: {patient JSON}
			//  Here's what you already found: [documents JSON]"
			promptBuilder.WriteString("\n")
		}
		promptBuilder.WriteString("Here is the text chunk:\n")

		for lineIndex, line := range window.WindowLines {
			lineNumber := window.StartIndex + lineIndex + 1
			// "#{lineNumber}: #{line}\n"
			promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", lineNumber, line))
		}

		analyzedDocuments = append(analyzedDocuments, models.AnalyzedDocument{
			Title: fmt.Sprintf("test_%d", window.StartIndex),
			// Just for testing purposes
			StartLine: 0,
			EndLine:   1,
		})
		fmt.Println(promptBuilder.String())

		response, err := aiService.Query(context.Background(), promptBuilder.String(), nil)

		if err != nil {
			return nil, nil, fmt.Errorf("AI query failed: %w", err)
		}

		fmt.Printf("OpenAI Response: %+v\n", response)
		// TODO: Send finalPrompt to OpenAI
		// TODO: Process response and append to analyzedDocuments

		fmt.Printf("OpenAI Response: %+v\n", response)

		// Extract and merge patient data from response
		// See Q&A 2025-10-14 for more info on this syntax
		if patientData, ok := response["patient"].(map[string]interface{}); ok {
			// Merge patient fields (prefer non-empty values)
			if name, ok := patientData["name"].(string); ok && name != "" {
				patient.Name = name
			}
			if species, ok := patientData["species"].(string); ok && species != "" {
				speciesPtr := species
				patient.Species = &speciesPtr
			}
			// ... continue for other fields
		}

		// Extract documents from response
		if docs, ok := response["documents"].([]interface{}); ok {
			for _, doc := range docs {
				if docMap, ok := doc.(map[string]interface{}); ok {
					// Check if this document already exists (deduplicate by start_line)
					startLine := int64(docMap["start_line"].(float64))

					// Check for duplicates
					isDuplicate := false
					for _, existingDoc := range analyzedDocuments {
						if existingDoc.StartLine == startLine {
							isDuplicate = true
							break
						}
					}

					if !isDuplicate {
						// Create new AnalyzedDocument and append
						analyzedDocuments = append(analyzedDocuments, models.AnalyzedDocument{
							Title:     docMap["title"].(string),
							StartLine: startLine,
							EndLine:   int64(docMap["end_line"].(float64)),
							// ... other fields
						})
					}
				}
			}
		}
	}

	return &patient, analyzedDocuments, nil
}
