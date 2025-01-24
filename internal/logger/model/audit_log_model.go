package model

import "time"

type AuditLog struct {
	TraceID   string                 `json:"trace_id" validate:"required"`
	Timestamp time.Time              `json:"timestamp" validate:"required"`
	Event     string                 `json:"event" validate:"required"`
	User      string                 `json:"user" validate:"required"`
	IpAddress string                 `json:"ip_address" validate:"required"`
	Status    string                 `json:"status" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}
