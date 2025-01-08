package request

import (
	"encoding/json"
	"mime/multipart"
)

type CreateProduct struct {
	Name        string                  `form:"name" binding:"required"`
	Price       json.Number             `form:"price" binding:"required"`
	Stock       json.Number             `form:"stock"`
	Description string                  `form:"description"`
	Categories  []int                   `form:"categories" binding:"required,min=1,dive,required"`
	Pictures    []*multipart.FileHeader `form:"pictures" binding:"required,filesize=2,filetype=image"`
}

type UpdateProduct struct {
	Name             string                  `form:"name" binding:"required"`
	Price            json.Number             `form:"price" binding:"required"`
	Stock            json.Number             `form:"stock"`
	Description      string                  `form:"description"`
	Categories       []int                   `form:"categories" binding:"required,min=1,dive,required"`
	ExistingPictures []string                `form:"existing_pictures" binding:"required_without=NewPictures"`
	NewPictures      []*multipart.FileHeader `form:"new_pictures" binding:"required_without=ExistingPictures"`
}
