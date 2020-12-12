package model

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	PictureId     int
	ImageUrl      string
	LargeImageUrl string
	Name          string
	Description   string
	Author        string
	Width         int
	Height        int
}
