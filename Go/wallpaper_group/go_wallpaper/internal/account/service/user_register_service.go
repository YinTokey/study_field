package service

import (
	"fmt"
	"go_wallpaper/internal/account/dao"
	"go_wallpaper/internal/account/model"
	"go_wallpaper/internal/account/serializer"
	"go_wallpaper/pkg"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	d := dao.NewUserDao(pkg.InstanceDB())

	var existNickName = d.CheckNickNameExist(service.Nickname)

	if existNickName {
		return &serializer.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	var existUserName = d.CheckUserNameExist(service.UserName)
	if existUserName {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {

	user := model.User{
		Nickname:       service.Nickname,
		UserName:       service.UserName,
		PasswordDigest: service.Password,
		Status:         model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		fmt.Println(err)
		return *err
	}

	d := dao.NewUserDao(pkg.InstanceDB())

	// 加密密码
	if err := d.SetPassword(user, service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户

	if err := d.CreateUser(&user); err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(&user)
}
