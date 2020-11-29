package service

import (
	//"Week02/serializer"
	"Week02/dao"
	"Week02/models"
)

type PictureService struct {
	dao *dao.PictureDao
}

func NewPictureService() *PictureService {
	return &PictureService{
		dao: dao.NewPictureDao(models.InstanceDB()),
	}

}

func (s *PictureService) AddPicture(data *models.Picture) error {
	return s.dao.AddPicture(data)
}

func (s *PictureService) GetAll() ([]models.Picture, error) {

	return s.dao.GetAll()
}

