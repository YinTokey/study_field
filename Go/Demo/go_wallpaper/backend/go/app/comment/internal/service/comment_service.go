package service

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_wallpaper/backend/go/internal/comment/dao"
	"go_wallpaper/backend/go/internal/comment/model"
	"go_wallpaper/backend/go/pkg"
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

	primaryId, _ := pkg.NewGuid()

	// 构建index
	index := &model.CommentIndex{
		PID:       primaryId,
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
		CommentId:   primaryId,
		AtMemberIds: "",
		Ip:          0,
		Platform:    0,
		Device:      "",
		Message:     content,
		Meta:        "",
	}
	// 插入content
	d.AppendContent(nContent)

	//生成response
	rsp := s.GenResponse(*index, nContent)
	fmt.Println(rsp)
	// response 写入 redis
	s.StoreResponseToCache(rsp)
}

func (s *CommentService) FetchComments(id string) ([]*model.CommentResponse, error) {

	result, _ := s.FetchCommentsFromCache(id)
	fmt.Println("直接从redis 获取数据 ", result)
	if len(result) == 0 {
		fmt.Println("从 db 获取数据 ", result)
		result, _ = s.FetchCommentsFromDB(id)
	}

	return result, nil
}

func (s *CommentService) FetchCommentsFromDB(id string) ([]*model.CommentResponse, error) {
	d := dao.NewCommentDao(pkg.InstanceDB())
	d.CreatePicTable()

	indeics, err := d.GetIndexs(id)
	if err != nil {
		fmt.Println("index getting error ", err)
	}

	var result []*model.CommentResponse

	for _, obj := range indeics {

		content, _ := d.GetContent(obj.PID)

		rsp := s.GenResponse(obj, content)

		result = append(result, rsp)

		// 写入redis
		s.StoreResponseToCache(rsp)
	}

	return result, nil
}

func (s *CommentService) StoreIndexToCache(index *model.CommentIndex) {

}

func (s *CommentService) StoreContentToCache(content *model.CommentContent) {

	key := fmt.Sprintf("%s", content.CommentId)

	buf, err := json.Marshal(content)
	if err != nil {
		fmt.Println("json.Marshal(content) 失败", err)
		return
	}

	// 0 表示不会过期
	pkg.RedisClient.Set(key, buf, 0)
}

func (s *CommentService) StoreResponseToCache(resp *model.CommentResponse) {

	key := fmt.Sprintf("%s_%v", resp.ObjId, resp.Floor)

	buf, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json.Marshal(resp) 失败", err)
		return
	}

	// 0 表示不会过期
	pkg.RedisClient.Set(key, buf, 0)
	fmt.Println("写入redis, key = ", key)

	// 使用sorted set
	//z := redis.Z{Score: float64(resp.Floor), Member: resp}
	//pkg.RedisClient.ZAdd(key, z)

}

func (s *CommentService) FetchCommentsFromCache(id string) ([]*model.CommentResponse, error) {

	keys := s.GetCacheKeys(id)

	var result []*model.CommentResponse

	for _, key := range keys {

		redis.Bytes(pkg.RedisClient.Get(key).Result())

		buf, err := redis.Bytes(pkg.RedisClient.Get(key).Result())

		if err != nil {
			fmt.Println(err)
			continue
		}

		rsp := &model.CommentResponse{}
		err = json.Unmarshal(buf, rsp)

		if err != nil {
			fmt.Println(err)
			continue
		}

		result = append(result, rsp)
	}

	return result, nil
}

// 获取匹配key
func (s *CommentService) GetCacheKeys(id string) []string {

	var cursor uint64
	var keys []string
	var err error
	for {
		//*扫描所有key，每次20条
		pattern := fmt.Sprintf("%s*", id)
		keys, cursor, err = pkg.RedisClient.Scan(cursor, pattern, 20).Result()
		fmt.Println("match keys start , err ", keys, err)
		if cursor == 0 {
			break
		}
	}

	return keys

}

func (s *CommentService) GenResponse(index model.CommentIndex, content *model.CommentContent) *model.CommentResponse {

	rsp := &model.CommentResponse{
		ObjId:    index.ObjId,
		ObjType:  index.ObjType,
		MemberId: index.MemberId,
		Root:     index.Root,
		Parent:   index.Parent,
		Floor:    index.Floor,
		Message:  content.Message,
	}

	return rsp
}
