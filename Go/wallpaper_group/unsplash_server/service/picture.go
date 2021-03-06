package service

import (
	"unsplash_server/dao"
	"unsplash_server/model"
)


// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
var unsplash_base_url string = "https://api.unsplash.com"

type PictureService struct {
}

func NewPictureService() PictureService {
	return PictureService{}
}

func (service *PictureService) Papular(page int, pageSize int) ([]model.Picture, error) {

	//job := unsplash.NewUnsplashJob()
	//job.FetchPics()
	d := dao.NewPictureDao(model.InstanceDB())

	return d.GetPictures(page, pageSize)
}