package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/service/grpc"
	acgProto "go_wallpaper/protos/acg_server"
	"net/http"
	"strconv"
)

var (
	acgCli acgProto.AcgService
)

func AcgInit() {
	service := grpc.NewService(
	//micro.Flags(cmn.CustomFlags...),
	)
	// 初始化， 解析命令行参数等
	service.Init()

	cli := service.Client()

	client.WithAddress("127.0.0.1:50051")

	// 初始化一个acg服务的客户端
	acgCli = acgProto.NewAcgService("go.micro.service.acg", cli)

	// 直接调用这个 ip 而不使用 discovery
	//client.WithAddress("127.0.0.1:50051")

}

func FetchAcgList(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println(page, pageSize, acgCli)

	client.WithAddress("127.0.0.1:50051")

	resp, err := acgCli.List(context.TODO(), &acgProto.ListRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	})
	client.WithAddress("127.0.0.1:50051")

	if err != nil {
		//log.Println(err.Error())
		fmt.Println("grpc client 请求失败 ", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
		"data": resp.Data,
	})

	//1、Dail连接
	// TODO: 地址临时写死
	//conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer conn.Close()
	//
	//client := acgPackage.NewAcgServiceClient(conn)
	//
	//page, _ := strconv.Atoi(c.Query("page"))
	//pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	//
	//request := &acgPackage.ListRequest{
	//	PageSize: int32(pageSize),
	//	Page:     int32(page),
	//}
	//
	//result, err := client.List(context.Background(), request)
	//
	//if err != nil {
	//	fmt.Println("grpc 请求错误", err)
	//}
	//
	//fmt.Println(result)
	//c.JSON(200, result)

	//client := NewUnPictureServiceClient(conn)
	//
	//request := &UnPictureRequest{Page: 1, PageSize: 10}
	//
	//result, err := client.GetUnPictureInfo(context.Background(), request)
	//
	//if err != nil {
	//	fmt.Println("grpc 请求错误", err)
	//}
	//
	//var list []*model.Picture
	//
	//for _, data := range result.Piclist {
	//	var pic = &model.Picture{}
	//	//pic.PictureId = data.PictureId
	//	//pic.ImageUrl = data.ImageUrl
	//	//pic.LargeImageUrl = data.LargeImageUrl
	//	//pic.Author = data.Author
	//	//pic.Width = float64(data.Width)
	//	//pic.Height = float64(data.Height)
	//	//pic.Likes = float64(data.Likes)
	//	pic.Name = data.Name
	//	//pic.Description = data.Description
	//	list = append(list, pic)
	//}
	//
	//c.JSON(200, list)
}

func FetchAcgRandom(c *gin.Context) {

}
