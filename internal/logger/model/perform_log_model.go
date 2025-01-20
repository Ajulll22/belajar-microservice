package model

import "time"

type PerformLog struct {
	Timestamp      time.Time              `json:"timestamp"`
	Service        string                 `json:"service"`
	Endpoint       string                 `json:"endpoint"`
	ResponseTimeMS int                    `json:"response_time_ms"`
	StatusCode     int                    `json:"status_code"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}
