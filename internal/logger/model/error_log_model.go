package model

import "time"

type ErrorLog struct {
	TraceID   string    `json:"trace_id" validate:"required"`
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Service   string    `json:"service" validate:"required"`
	Error     struct {
		Message  string `json:"message" validate:"required"`
		Filename string `json:"filename" validate:"required"`
		Line     int    `json:"line"`
	} `json:"error" validate:"required"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
