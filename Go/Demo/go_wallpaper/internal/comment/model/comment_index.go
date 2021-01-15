package model

import (
	"github.com/jinzhu/gorm"
)

type CommentIndex struct {
	gorm.Model
	ObjId      int64   `json:"obj_id"`
	ObjType    int8    `json:"obj_type"`
	MemberId   int64   `json:"member_id"`
	Root       string  `json:"root"`
	Parent     string  `json:"parent"`
	Floor      string  `json:"floor"`
	Count      float64 `json:"count"`
	RootCount  float64 `json:"root_count"`
	Like       float64 `json:"like"`
	Hate       float64 `json:"hate"`
	State      float64 `json:"state"`
	Attrs      float64 `json:"attrs"`
	UpdateTime float64 `json:"update_time"`
	CreateTime float64 `json:"create_time"`
}
