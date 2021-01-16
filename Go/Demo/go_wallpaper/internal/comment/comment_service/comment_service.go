package comment_service

import (
	"fmt"
	"github.com/satori/go.uuid"
	"go_wallpaper/internal/comment/dao"
	"go_wallpaper/internal/comment/model"
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

	objId := id

	// 定义评论类型
	var objType int8 = 1

	sub, err := d.QueryCommentSubject(objId)
	if err != nil {
		fmt.Println("没有subject")
		// 构建subject
		subject := &model.CommentSubject{
			ObjId:     objId,
			ObjType:   objType,
			MemberId:  0,
			Count:     1,
			RootCount: 1,
			AllCount:  0,
			State:     0,
			Attrs:     0,
		}

		d.AppendSubject(subject)
	}

	// TODO: 更新subject
	fmt.Println(sub)

	// 获取原评论列表
	idxs, err := d.GetIndexs(objId)

	count := int32(len(idxs)) + 1

	fmt.Println("count .. ", count)

	idxObjId := uuid.Must(uuid.NewV4(), err).String()

	// 构建index
	index := &model.CommentIndex{
		ObjId:     idxObjId,
		ObjType:   objType,
		MemberId:  0,
		Root:      0,
		Parent:    0,
		Floor:     count,
		Count:     count,
		RootCount: count,
		Like:      0,
		Hate:      0,
		State:     0,
		Attrs:     0,
	}

	// 插入index
	err = d.AppendIndex(index)
	if err != nil {
		fmt.Println("插入index 失败")
	}

	// 构建content
	nContent := &model.CommentContent{
		CommentId:   idxObjId,
		AtMemberIds: "",
		Ip:          0,
		Platform:    0,
		Device:      "",
		Message:     content,
		Meta:        "",
	}

	d.AppendContent(nContent)

}
