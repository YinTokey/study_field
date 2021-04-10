package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"unsplash_server/internal/service"
	proto "unsplash_server/proto"
)

// UnPicture : 用于实现UnPictureServiceServiceHandler接口的对象
type UnPicture struct{}

func FetchPapular(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println("page %s , page sie %s", page, pageSize)

	fetchFromServcie(c, page, pageSize)
}

func fetchFromServcie(c *gin.Context, page int, pageSize int) {
	service := service.NewPictureService()

	if err := c.ShouldBind(&service); err == nil {
		res, err := service.Papular(page, pageSize)
		if err != nil {
			c.JSON(200, "error happend")
		} else {
			c.JSON(200, res)
		}

	} else {
		c.JSON(200, err)
	}

}

// 该方法用于实现 go-micro 生成的接口
func (a *UnPicture) GetUnPictureInfo(ctx context.Context, req *proto.UnPictureRequest, res *proto.UnPictureInfo) error {

	fmt.Println(" Register -- grpc -- imp")

	page := int(req.Page)
	pageSize := int(req.PageSize)

	service := service.NewPictureService()
	serviceRes, err := service.Papular(page, pageSize)

	if err != nil {
		return err
	}

	var list []*proto.UnPictureInfo_Picture

	for _, data := range serviceRes {
		pic := &proto.UnPictureInfo_Picture {
			PictureId: data.PictureId,
			ImageUrl: data.ImageUrl,
			LargeImageUrl: data.LargeImageUrl,
			Author: data.Author,
			Width: int32(data.Width),
			Height: int32(data.Height),
			Likes: data.Likes,
			Name: data.Name,
			Description: data.Description,
			Tags: data.Tags,
		}

		list = append(list, pic)

	}
	res.Piclist = list

	return nil
}


