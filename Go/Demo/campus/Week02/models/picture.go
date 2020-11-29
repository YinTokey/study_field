package model

import (
	"github.com/jinzhu/gorm"

)

// Picture 模型
type Picture struct {
	gorm.Model
	URL       	string
	Author      string
	About       string
	UserID		string
}

