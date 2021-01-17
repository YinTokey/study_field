package model

import (
	"github.com/jinzhu/gorm"
)

// 表示评论一层楼
type CommentIndex struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey"`
	ObjId     string `json:"obj_id"`
	ObjType   int8   `json:"obj_type"`
	MemberId  int64  `json:"member_id"`
	Root      int64  `json:"root"`
	Parent    int64  `json:"parent"`
	Floor     int32  `json:"floor"`
	Count     int32  `json:"count"`
	RootCount int32  `json:"root_count"`
	Like      int32  `json:"like"`
	Hate      int32  `json:"hate"`
	State     int8   `json:"state"`
	Attrs     int32  `json:"attrs"`
}
