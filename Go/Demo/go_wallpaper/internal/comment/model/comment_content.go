package model

import (
	"github.com/jinzhu/gorm"
)

type CommentContent struct {
	gorm.Model
	CommentId   string  `json:"comment_id"`
	AtMemberIds string  `json:"at_member_ids"`
	Ip          string  `json:"ip"`
	Platform    string  `json:"platform"`
	Device      string  `json:"device"`
	Message     string  `json:"message"`
	Meta        float64 `json:"meta"`
	UpdateTime  float64 `json:"update_time"`
	CreateTime  float64 `json:"create_time"`
}
