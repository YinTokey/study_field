package model

// 组装完返回的单条评论
type CommentResponse struct {
	ObjId    string `json:"obj_id"`
	ObjType  int8   `json:"obj_type"`
	MemberId int64  `json:"member_id"`
	Root     int64  `json:"root"`
	Parent   int64  `json:"parent"`
	Floor    int32  `json:"floor"`
	Message  string `json:"message"`
}
