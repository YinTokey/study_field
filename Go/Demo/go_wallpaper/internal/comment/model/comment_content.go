package model

import (
	"github.com/jinzhu/gorm"
)

type CommentContent struct {
	gorm.Model
	CommentId   int64  `json:"comment_id"`
	AtMemberIds string `json:"at_member_ids"`
	Ip          int64  `json:"ip"`
	Platform    int8   `json:"platform"`
	Device      string `json:"device"`
	Message     string `json:"message"`
	Meta        string `json:"meta"`
}
