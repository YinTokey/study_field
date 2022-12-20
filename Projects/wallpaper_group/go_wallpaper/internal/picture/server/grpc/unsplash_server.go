package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	//"go_wallpaper/configs"
	//"go_wallpaper/internal/picture/service"
	//"google.golang.org/grpc"
	//"net"
)

type UnPictureServiceImpl struct {
}

//func (us *UnPictureServiceImpl) GetUnPictureInfo(ctx context.Context, request *api.UnPictureRequest) (*api.UnPictureInfo, error) {
//page := request.Page
//pageSize := request.PageSize
//srv := service.NewUmCollectService()
//
//fmt.Println("GetUnPictureInfo comment_service start")
//
//result, err := srv.Papular(int(page), int(pageSize))
//if err != nil {
//	return nil, err
//}
//
//var list []*api.UnPictureInfo_Picture
//
//for _, data := range result {
//	var pic api.UnPictureInfo_Picture = api.UnPictureInfo_Picture{}
//	pic.PictureId = data.PictureId
//	pic.ImageUrl = data.ImageUrl
//	pic.LargeImageUrl = data.LargeImageUrl
//	pic.Author = data.Author
//	pic.Width = float32(data.Width)
//	pic.Height = float32(data.Height)
//	pic.Likes = float32(data.Likes)
//	pic.Name = data.Name
//	pic.Description = data.Description
//
//	list = append(list, &pic)
//}
//
//response := &api.UnPictureInfo{
//	Piclist: list,
//}

//return response, nil

//	return nil, nil
//}

func Serve() error {

	fmt.Println("grpc 服务启动中。。")

	server := grpc.NewServer()

	api.RegisterUnPictureServiceServer(server, new(UnPictureServiceImpl))

	lis, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("监听错误")
		panic(err.Error())
		return err
	}
	server.Serve(lis)

	fmt.Println("grpc 服务启动完成")

	return nil
}

func main() {

	fmt.Println("grpc 服务启动中。。")

	//configs.Init()
	//
	//server := grpc.NewServer()
	//
	//api.RegisterUnPictureServiceServer(server, new(UnPictureServiceImpl))
	//
	//lis, err := net.Listen("tcp", ":8090")
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println("grpc 服务启动开始")
	//
	//server.Serve(lis)

}
