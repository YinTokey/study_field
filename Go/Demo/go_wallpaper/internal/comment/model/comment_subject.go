package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
)

type CommentSubject struct {
	gorm.Model
	ObjId      int64               `json:"obj_id"`
	ObjType    int8                `json:"obj_type"`
	MemberId   int64               `json:"member_id"`
	Count      int32               `json:"count"`
	RootCount  int32               `json:"root_count"`
	AllCount   int32               `json:"all_count"`
	State      int8                `json:"state"`
	Attrs      int32               `json:"attrs"`
	UpdateTime timestamp.Timestamp `json:"update_time"`
	CreateTime timestamp.Timestamp `json:"create_time"`
}
