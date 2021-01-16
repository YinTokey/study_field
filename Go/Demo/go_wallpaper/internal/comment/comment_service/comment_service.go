package comment_service

import (
	"go_wallpaper/internal/comment/dao"
	"go_wallpaper/pkg"
)

type CommentService struct {
}

func NewCommentService() CommentService {
	return CommentService{}
}

func (s *CommentService) AddComment(id string, content string) {

	d := dao.NewCommentDao(pkg.InstanceDB())
	d.CreatePicTable()
}
