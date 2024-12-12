package model

import "mime/multipart"

type Media struct {
	ID   string                `json:"id"`
	File *multipart.FileHeader `json:"-"`
}
