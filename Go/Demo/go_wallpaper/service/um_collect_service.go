package service

import (
	"go_wallpaper/model"
)

// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
var unsplash_base_url string = "https://api.unsplash.com"

type UmCollectService struct {
}

func NewUmCollectService() UmCollectService {
	return UmCollectService{}
}

func (service *UmCollectService) Papular() []model.Picture {

	//job := unsplash.NewUnsplashJob()
	//job.FetchPics()

	return nil
}
