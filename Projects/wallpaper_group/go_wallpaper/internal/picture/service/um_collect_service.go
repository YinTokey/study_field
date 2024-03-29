package service

import (
	"go_wallpaper/internal/picture/dao"
	"go_wallpaper/internal/picture/model"
	"go_wallpaper/pkg"
)

// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
var unsplash_base_url string = "https://api.unsplash.com"

type UmCollectService struct {
}

func NewUmCollectService() UmCollectService {
	return UmCollectService{}
}

func (service *UmCollectService) Papular(page int, pageSize int) ([]model.Picture, error) {

	//job := unsplash.NewUnsplashJob()
	//job.FetchPics()
	d := dao.NewPictureDao(pkg.InstanceDB())

	return d.GetPictures(page, pageSize)
}
