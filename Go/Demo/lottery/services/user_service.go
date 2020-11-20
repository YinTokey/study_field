package services

import (
	"fmt"
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type UserService interface {
	GetAll(page, size int) []models.LtUser
	CountAll() int

	Get(id int) *models.LtUser
	Update(user *models.LtUser, colums []string) error
	Create(user *models.LtUser) error
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}

func (s *userService) Get(id int) *models.LtUser {
	data := s.dao.Get(id)
	return data
}

func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}

func (s *userService) CountAll() int {
	return s.dao.CountAll()
}

func (s *userService) Create(data *models.LtUser) error {
	fmt.Println("创建用户")
	return s.dao.Create(data)
}

func (s *userService) Update(user *models.LtUser, colums []string) error {
	return nil
}