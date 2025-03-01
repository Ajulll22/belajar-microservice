package request

import "mime/multipart"

type UploadMedia struct {
	Files []*multipart.FileHeader `form:"files" binding:"required,filesize=2"`
}

type DeleteMedia struct {
	ID string `json:"id"`
}
