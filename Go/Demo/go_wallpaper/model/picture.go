package model

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	PictureId     string
	ImageUrl      string
	LargeImageUrl string
	Name          string
	Description   string
	Author        string
	//UpdateAt      time.Time
	Width      float64
	Height     float64
	Likes      float64
	Categories []string
}
