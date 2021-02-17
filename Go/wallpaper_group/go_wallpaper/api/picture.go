package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wallpaper/internal/comment/comment_service"
	"go_wallpaper/internal/picture/service"
	timestamppb "go_wallpaper/protos/acg_server"
	"google.golang.org/grpc"
	"strconv"
)

var (
// unsplashCli unsplashProto.UnPictureService
)

func PictureInit() {
	//service := grpc.NewService(
	////micro.Flags(cmn.CustomFlags...),
	//)
	//// 初始化， 解析命令行参数等
	//service.Init()
	//
	//cli := service.Client()
	//
	//// 初始化一个account服务的客户端
	//unsplashCli = unsplashProto.NewUnPictureService("go.micro.service.unsplash", cli)
	//// 初始化一个upload服务的客户端

}

func Fetch500pxPapular(c *gin.Context) {
	//fmt.Println("500px fetch .....")
	//var comment_service comment_service.PxCollectService = comment_service.NewPxCollectService()
	fmt.Println("unsplash fetch ...")

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println("page %s , page sie %s", page, pageSize)

	// 基于grpc 获取
	fetchFromGRPC(c, page, pageSize)

	// 基于单体架构 comment_service 获取
	//fetchFromServcie(c, page, pageSize)
}

func Fetch500pxDetail(c *gin.Context) {

	fmt.Println("fetch  detail ...")

	objId := c.Query("id")

	service := comment_service.NewCommentService()

	res, err := service.FetchComments(objId)
	fmt.Println("返回给前端的结果 ", res)
	if err != nil {
		c.JSON(200, "error happend")
	} else {
		c.JSON(200, res)
	}
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
	// 基于原始 grpc 服务调用

	target := "127.0.0.1:10087"

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := timestamppb.NewAcgServiceClient(conn)

	request := &timestamppb.ListRequest{
		PageSize: int32(pageSize),
		Page:     int32(page),
	}

	result, err := client.List(context.Background(), request)

	if err != nil {
		fmt.Println("grpc 请求错误", err)
	}

	//fmt.Println(result)
	c.JSON(200, result)

	// 基于 Go micro 服务调用
	//resp, err := unsplashCli.GetUnPictureInfo(context.TODO(), &unsplashProto.UnPictureRequest{
	//	Page:     int64(page),
	//	PageSize: int64(pageSize),
	//})
	//
	//if err != nil {
	//	log.Println(err.Error())
	//	c.Status(http.StatusInternalServerError)
	//	return
	//}
	//
	//c.JSON(200, resp)

}
