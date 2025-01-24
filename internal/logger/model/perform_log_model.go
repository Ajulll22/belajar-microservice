package model

import "time"

type PerformLog struct {
	TraceID        string                 `json:"trace_id" validate:"required"`
	Timestamp      time.Time              `json:"timestamp" validate:"required"`
	Service        string                 `json:"service" validate:"required"`
	Endpoint       string                 `json:"endpoint" validate:"required"`
	ResponseTimeMS int                    `json:"response_time_ms" validate:"required"`
	StatusCode     int                    `json:"status_code" validate:"required"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}
