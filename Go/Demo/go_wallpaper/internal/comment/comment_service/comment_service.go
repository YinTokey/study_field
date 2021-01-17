package comment_service

import (
	"fmt"
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

	fmt.Println("count .. ", len(idxs))

	// 构建index
	index := &model.CommentIndex{
		ObjId:     objId,
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
		AtMemberIds: "",
		Ip:          0,
		Platform:    0,
		Device:      "",
		Message:     content,
		Meta:        "",
	}

	d.AppendContent(nContent)

}

func (s *CommentService) FetchComments(id string) ([]model.CommentResponse, error) {

	result, _ := s.FetchCommentsFromCache(id)

	if len(result) == 0 {
		result, _ = s.FetchCommentsFromDB(id)
	}

	return result, nil
}

func (s *CommentService) FetchCommentsFromDB(id string) ([]model.CommentResponse, error) {
	d := dao.NewCommentDao(pkg.InstanceDB())
	d.CreatePicTable()

	indeics, err := d.GetIndexs(id)
	if err != nil {
		fmt.Println("index getting error ", err)
	}

	var result []model.CommentResponse

	for _, obj := range indeics {

		//content, _ := d.GetContent(obj.ID)

		content, _ := d.GetContent(1)

		rsp := model.CommentResponse{
			ObjId:    obj.ObjId,
			ObjType:  obj.ObjType,
			MemberId: obj.MemberId,
			Root:     obj.Root,
			Parent:   obj.Parent,
			Floor:    obj.Floor,
			Message:  content.Message,
		}

		result = append(result, rsp)
	}

	return result, nil
}

func (s *CommentService) FetchCommentsFromCache(id string) ([]model.CommentResponse, error) {
	return nil, nil
}
