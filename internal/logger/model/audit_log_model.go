package model

import "time"

type AuditLog struct {
	Timestamp time.Time              `json:"timestamp"`
	Event     string                 `json:"event"`
	User      string                 `json:"user"`
	IpAddress string                 `json:"ip_address"`
	Status    string                 `json:"status"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}
