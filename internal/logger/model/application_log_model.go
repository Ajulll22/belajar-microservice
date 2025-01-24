package model

import "time"

type ApplicationLog struct {
	TraceID   string                 `json:"trace_id" validate:"required"`
	Timestamp time.Time              `json:"timestamp" validate:"required"`
	Level     string                 `json:"level" validate:"required"`
	Service   string                 `json:"service" validate:"required"`
	Message   string                 `json:"message" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}
