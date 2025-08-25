package models

import "time"

type UnprocessedDocument struct {
	Content       string    `json:"content" db:"content"`
	NumberOfLines int64     `json:"numberOfLines" db:"num_lines"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}
