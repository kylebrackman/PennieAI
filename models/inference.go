package models

import (
	"time"
)

type Inference struct {
	ID            int64                  `json:"id" db:"id"`
	Request       string                 `json:"request" db:"request"`
	Response      string                 `json:"response" db:"response"`
	Config        map[string]interface{} `json:"config" db:"config"`
	InferableType *string                `json:"inferableType" db:"inferable_type"`
	InferableID   *int64                 `json:"inferableId" db:"inferable_id"`
	CreatedAt     time.Time              `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time              `json:"updatedAt" db:"updated_at"`
}
