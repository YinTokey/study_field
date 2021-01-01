package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_wallpaper/internal/account/serializer"
	"go_wallpaper/internal/account/service"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {

	service := service.UserRegisterService{
		UserName:        c.PostForm("user_name"),
		Password:        c.PostForm("password"),
		PasswordConfirm: c.PostForm("password_confirm"),
		Nickname:        c.PostForm("nickname"),
	}

	fmt.Println("start register")
	res := service.Register()
	c.JSON(200, res)

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		//c.JSON(200, serializer.ErrorResponse(err))
		c.JSON(200, err)
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	//var user *model.User = middleware.CurrentUser(c)
	//res := serializer.BuildUserResponse(user)
	//c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
