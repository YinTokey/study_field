package model

import "github.com/jinzhu/gorm"

type Page struct {
	gorm.Model
	CurrentPage int    `json:"current_page"`
	TotalPages  int    `json:"total_pages"`
	TotalItems  int    `json:"total_items"`
	Feature     string `json:"feature"`
	Filters     struct {
	} `json:"filters"`
	Photos      []Photo `json:"photos"`

}
