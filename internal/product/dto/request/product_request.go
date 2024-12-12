package request

import (
	"encoding/json"
	"mime/multipart"
)

type ProductInsert struct {
	Name        string                  `form:"name" binding:"required"`
	Price       json.Number             `form:"price"`
	Stock       json.Number             `form:"stock"`
	Description string                  `form:"description"`
	Pictures    []*multipart.FileHeader `form:"pictures" binding:"required,filesize=2,filetype=image"`
	Categories  []int                   `form:"categories" binding:"required,min=1,dive,required"`
}
