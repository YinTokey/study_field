package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wallpaper/internal/comment/comment_service"
	"go_wallpaper/internal/picture/model"
	"go_wallpaper/internal/picture/service"

	"google.golang.org/grpc"
	"strconv"
)

// UserRegister 用户注册接口
func Fetch500pxPapular(c *gin.Context) {
	//fmt.Println("500px fetch .....")
	//var comment_service comment_service.PxCollectService = comment_service.NewPxCollectService()
	fmt.Println("unsplash fetch ...")

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println("page %s , page sie %s", page, pageSize)

	// 基于grpc 获取
	//fetchFromGRPC(c, page, pageSize)

	// 基于单体架构 comment_service 获取
	fetchFromServcie(c, page, pageSize)
}

func Fetch500pxDetail(c *gin.Context) {

	fmt.Println("fetch  detail ...")

	pictureId := c.Query("id")
	fmt.Println(pictureId)

}

func PostComment(c *gin.Context) {
	fmt.Println("post comment ...")

	obj_id := c.PostForm("id")
	content := c.PostForm("content")

	fmt.Println(obj_id, "\n", content)

	service := comment_service.NewCommentService()

	service.AddComment(obj_id, content)

}

func fetchFromServcie(c *gin.Context, page int, pageSize int) {
	service := service.NewUmCollectService()

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

func fetchFromGRPC(c *gin.Context, page int, pageSize int) {

	//1、Dail连接
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	client := NewUnPictureServiceClient(conn)

	request := &UnPictureRequest{Page: 1, PageSize: 10}

	result, err := client.GetUnPictureInfo(context.Background(), request)

	if err != nil {
		fmt.Println("grpc 请求错误", err)
	}

	var list []*model.Picture

	for _, data := range result.Piclist {
		var pic = &model.Picture{}
		pic.PictureId = data.PictureId
		pic.ImageUrl = data.ImageUrl
		pic.LargeImageUrl = data.LargeImageUrl
		pic.Author = data.Author
		pic.Width = float64(data.Width)
		pic.Height = float64(data.Height)
		pic.Likes = float64(data.Likes)
		pic.Name = data.Name
		pic.Description = data.Description
		list = append(list, pic)
	}

	c.JSON(200, list)

}
