package models

import (
	"time"
)

type Inference struct {
	ID            int64                  `json:"id" repository:"id"`
	Request       string                 `json:"request" repository:"request"`
	Response      string                 `json:"response" repository:"response"`
	Config        map[string]interface{} `json:"config" repository:"config"`
	InferableType *string                `json:"inferableType" repository:"inferable_type"`
	InferableID   *int64                 `json:"inferableId" repository:"inferable_id"`
	CreatedAt     time.Time              `json:"createdAt" repository:"created_at"`
	UpdatedAt     time.Time              `json:"updatedAt" repository:"updated_at"`
}
