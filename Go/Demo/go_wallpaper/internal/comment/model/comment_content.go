package model

import (
	"github.com/jinzhu/gorm"
)

// 单条评论(一级，二级)
type CommentContent struct {
	gorm.Model
	CommentId   string `gorm:"primaryKey"`
	AtMemberIds string `json:"at_member_ids"`
	Ip          int64  `json:"ip"`
	Platform    int8   `json:"platform"`
	Device      string `json:"device"`
	Message     string `json:"message"`
	Meta        string `json:"meta"`
}
