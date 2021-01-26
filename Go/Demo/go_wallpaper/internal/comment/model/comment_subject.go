package model

import (
	"github.com/jinzhu/gorm"
)

// 整个评论列表概述
type CommentSubject struct {
	gorm.Model
	ObjId     string `json:"obj_id"`
	ObjType   int8   `json:"obj_type"`
	MemberId  int64  `json:"member_id"`
	Count     int32  `json:"count"`
	RootCount int32  `json:"root_count"`
	AllCount  int32  `json:"all_count"`
	State     int8   `json:"state"`
	Attrs     int32  `json:"attrs"`
}
