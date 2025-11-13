package models

import "time"

type AnalyzedDocument struct {
	ID                    int64     `json:"id" db:"id"`
	Title                 string    `json:"title" db:"title"`
	Content               string    `json:"content" db:"content"`
	NumberOfLines         int64     `json:"numberOfLines" db:"num_lines"`
	PatientID             int64     `json:"patientId" db:"patient_id"`
	StartLine             int64     `json:"startLine" db:"start_line"`
	EndLine               int64     `json:"endLine" db:"end_line"`
	UnprocessedDocumentId int64     `json:"unprocessedDocumentId" db:"unprocessed_document_id"`
	CreatedAt             time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt             time.Time `json:"updatedAt" db:"updated_at"`
	WindowLines           []string  `json:"windowLines" db:"window_lines"`
}
