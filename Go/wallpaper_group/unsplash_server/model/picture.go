
package model

import (
"github.com/jinzhu/gorm"
)

type Picture struct {
	gorm.Model
	PictureId     int64  `json:"picture_id"`
	ImageUrl      string `json:"image_url"`
	LargeImageUrl string `json:"large_image_url"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Likes         int32  `json:"likes"`
	Tags    	  string `json:"categories"`
}