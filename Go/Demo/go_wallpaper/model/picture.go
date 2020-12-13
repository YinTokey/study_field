package model

import (
	"github.com/jinzhu/gorm"
)

type Picture struct {
	gorm.Model
	PictureId     string  `json:"picture_id"`
	ImageUrl      string  `json:"image_url"`
	LargeImageUrl string  `json:"large_image_url"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Author        string  `json:"author"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	Likes         float64 `json:"likes"`
	//Categories []string `json:"categories"`
}
