package constant

import "time"

const (
	CacheTTLOneDay     = 24 * time.Hour
	CacheTTLFiveMinute = 5 * time.Minute
	CacheTTLForever    = 0
	CacheTTLInvalidate = -1
)
