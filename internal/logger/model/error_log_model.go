package model

import "time"

type ErrorLog struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Service   string    `json:"service"`
	Error     struct {
		Message  string `json:"message"`
		Filename string `json:"filename"`
		Line     int    `json:"line"`
	} `json:"error"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
