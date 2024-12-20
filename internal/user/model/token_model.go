package model

import "time"

type TokenDetails struct {
	AccessToken   string
	RefreshToken  string
	AccessUUID    string
	RefreshUUID   string
	AccessExpiry  time.Time
	RefreshExpiry time.Time
}
