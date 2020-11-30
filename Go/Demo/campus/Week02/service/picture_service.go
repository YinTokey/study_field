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

func (s *PictureService) Query(id int) (string, error) {
	return s.dao.Query(id)
}

