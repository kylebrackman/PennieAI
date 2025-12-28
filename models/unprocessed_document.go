package models

import "time"

type UnprocessedDocument struct {
	ID            int64     `json:"id" repository:"id"`
	Content       string    `json:"content" repository:"content"`
	NumberOfLines int64     `json:"numberOfLines" repository:"num_lines"`
	CreatedAt     time.Time `json:"createdAt" repository:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" repository:"updated_at"`
}
