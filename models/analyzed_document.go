package models

import "time"

type AnalyzedDocument struct {
	ID                    int64     `json:"id" repository:"id"`
	Title                 string    `json:"title" repository:"title"`
	Content               string    `json:"content" repository:"content"`
	NumberOfLines         int64     `json:"numberOfLines" repository:"num_lines"`
	PatientID             int64     `json:"patientId" repository:"patient_id"`
	StartLine             int64     `json:"startLine" repository:"start_line"`
	EndLine               int64     `json:"endLine" repository:"end_line"`
	UnprocessedDocumentId int64     `json:"unprocessedDocumentId" repository:"unprocessed_document_id"`
	CreatedAt             time.Time `json:"createdAt" repository:"created_at"`
	UpdatedAt             time.Time `json:"updatedAt" repository:"updated_at"`
	WindowLines           []string  `json:"windowLines" repository:"window_lines"`
}
