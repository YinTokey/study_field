package api

import (
	"context"
	"unsplash_server/internal/service"
	proto "unsplash_server/proto"
)

type UnsplashServer struct{
	proto.UnimplementedUnPictureServiceServer
}

func (u *UnsplashServer) GetUnPictureInfo(ctx context.Context, req *proto.UnPictureRequest) (*proto.UnPictureInfo, error) {
	page := int(req.Page)
	pageSize := int(req.PageSize)

	service := service.NewPictureService()
	serviceRes, err := service.Papular(page, pageSize)

	if err != nil {
		return nil,err
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

	res := &proto.UnPictureInfo{}

	res.Piclist = list

	return res,nil
}

