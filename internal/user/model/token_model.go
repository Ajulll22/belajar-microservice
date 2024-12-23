package model

import "time"

type TokenDetails struct {
	UserID        int
	AccessToken   string
	RefreshToken  string
	AccessUUID    string
	RefreshUUID   string
	AccessExpiry  time.Time
	RefreshExpiry time.Time
}
