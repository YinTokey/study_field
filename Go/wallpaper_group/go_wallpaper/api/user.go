package api

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/service/grpc"
	"go_wallpaper/internal/account/serializer"
	accountProto "go_wallpaper/protos/account_server"
	"log"
	"net/http"
)

var (
	accountCli accountProto.AccountService
)

func AccountInit() {
	service := grpc.NewService(
	//micro.Flags(cmn.CustomFlags...),
	)
	// 初始化， 解析命令行参数等
	service.Init()

	cli := service.Client()
	// tracer, err := tracing.Init("apigw service", "<jaeger-agent-host>")
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	cli = client.NewClient(
	// 		client.Wrap(mopentracing.NewClientWrapper(tracer)),
	// 	)
	// }

	// 初始化一个account服务的客户端
	accountCli = accountProto.NewAccountService("go.micro.service.account", cli)
	// 初始化一个upload服务的客户端

}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {

	//service := service.UserRegisterService{
	//	UserName:        c.PostForm("user_name"),
	//	Password:        c.PostForm("password"),
	//	PasswordConfirm: c.PostForm("password_confirm"),
	//	Nickname:        c.PostForm("nickname"),
	//}
	//
	//fmt.Println("start register")
	//res := service.Register()
	//c.JSON(200, res)

	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	passwordConfirm := c.PostForm("password_confirm")
	nickname := c.PostForm("nickname")

	resp, err := accountCli.Register(context.TODO(), &accountProto.RegisterRequest{
		Username:        userName,
		Nickname:        nickname,
		Password:        password,
		PasswordConfirm: passwordConfirm,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {

	//service := service.UserLoginService{
	//	UserName: c.PostForm("user_name"),
	//	Password: c.PostForm("password"),
	//}
	//res := service.Login(c)
	//c.JSON(200, res)

	userName := c.PostForm("user_name")
	password := c.PostForm("password")

	resp, err := accountCli.Login(context.TODO(), &accountProto.LoginRequest{
		Username: userName,
		Password: password,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	//var user *model.User = middleware.CurrentUser(c)
	//res := serializer.BuildUserResponse(user)
	//c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	fmt.Println("用户登出")
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
