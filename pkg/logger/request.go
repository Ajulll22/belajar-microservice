package logger

import "time"

type ApplicationLog struct {
	TraceID   string                 `json:"trace_id" validate:"required"`
	Timestamp time.Time              `json:"timestamp" validate:"required"`
	Level     string                 `json:"level" validate:"required"`
	Service   string                 `json:"service" validate:"required"`
	Message   string                 `json:"message" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type AuditLog struct {
	TraceID   string                 `json:"trace_id" validate:"required"`
	Timestamp time.Time              `json:"timestamp" validate:"required"`
	Event     string                 `json:"event" validate:"required"`
	User      string                 `json:"user" validate:"required"`
	IpAddress string                 `json:"ip_address" validate:"required"`
	Status    string                 `json:"status" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

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

type PerformLog struct {
	TraceID        string                 `json:"trace_id" validate:"required"`
	Timestamp      time.Time              `json:"timestamp" validate:"required"`
	Service        string                 `json:"service" validate:"required"`
	Endpoint       string                 `json:"endpoint" validate:"required"`
	ResponseTimeMS int                    `json:"response_time_ms" validate:"required"`
	StatusCode     int                    `json:"status_code" validate:"required"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}
