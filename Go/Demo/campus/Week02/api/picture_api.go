package api

import (
	"Week02/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddPicture(c *gin.Context) {
	service := service.NewPictureService()
	err := c.ShouldBind(&service)
	if err != nil {
		//c.JSON(200, ErrorResponse(err))
		return
	}
	err = service.AddPicture(nil)
	if err != nil {
		// 查询错误
		fmt.Println(err)
		return
	}
	//c.JSON(200, res)

}

func QueryPicture(c *gin.Context) {
	service := service.NewPictureService()
	err := c.ShouldBind(&service)
	if err != nil {
		//c.JSON(200, ErrorResponse(err))
		return;
	}
	res, err := service.GetAll()
	if err != nil {
		// 查询错误
		return;
	}
	c.JSON(200, res)

}


//// UserRegister 用户注册接口
//func UserRegister(c *gin.Context) {
//	var service service.UserRegisterService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Register()
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//// UserLogin 用户登录接口
//func UserLogin(c *gin.Context) {
//	var service service.UserLoginService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Login(c)
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//// UserMe 用户详情
//func UserMe(c *gin.Context) {
//	user := CurrentUser(c)
//	res := serializer.BuildUserResponse(*user)
//	c.JSON(200, res)
//}
//
//// UserLogout 用户登出
//func UserLogout(c *gin.Context) {
//	s := sessions.Default(c)
//	s.Clear()
//	s.Save()
//	c.JSON(200, serializer.Response{
//		Code: 0,
//		Msg:  "登出成功",
//	})
//}
