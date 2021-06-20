package service

import (
	"unsplash_server/global"
	"unsplash_server/internal/dao"
	"unsplash_server/internal/model"
)

// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
// var unsplash_base_url string = "https://api.unsplash.com"

type PictureService struct {
}

func NewPictureService() PictureService {
	return PictureService{}
}

func (service *PictureService) Papular(page int, pageSize int) ([]model.Picture, error) {

	//job := unsplash.NewUnsplashJob()
	//job.FetchPics()
	d := dao.NewPictureDao(global.DBEngine)

	return d.GetPictures(page, pageSize)
}
