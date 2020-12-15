package server

import (
	"context"
	"go_wallpaper/message"
	"go_wallpaper/service"
	"google.golang.org/grpc"
	"net"
)

type UnPictureServiceImpl struct {
}

//func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {

func (us *UnPictureServiceImpl) GetUnPictureInfo(ctx context.Context, request *message.UnPictureRequest) (*message.UnPictureInfo, error) {
	page := request.Page
	pageSize := request.PageSize
	service := service.NewUmCollectService()

	result, err := service.Papular(int(page), int(pageSize))
	if err != nil {
		return nil, err
	}

	var list []*message.UnPictureInfo_Picture

	for _, data := range result {
		var pic message.UnPictureInfo_Picture = message.UnPictureInfo_Picture{}
		pic.PictureId = data.PictureId
		pic.ImageUrl = data.ImageUrl
		pic.LargeImageUrl = data.LargeImageUrl
		pic.Author = data.Author
		pic.Width = float32(data.Width)
		pic.Height = float32(data.Height)
		pic.Likes = float32(data.Likes)
		pic.Name = data.Name
		pic.Description = data.Description

		list = append(list, &pic)
	}

	response := &message.UnPictureInfo{
		Piclist: list,
	}

	return response, nil
}

func main() {

	server := grpc.NewServer()

	message.RegisterUnPictureServiceServer(server, new(UnPictureServiceImpl))

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)

}
