// gprc 请求实现

package api

import (
	"time"
	"context"
	"fmt"
	proto "account_server/proto"
	"account_server/util"
	"account_server/service"
)

// Account : 用于实现AcountServiceHandler接口的对象
type Account struct{}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

func (a *Account) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {

	fmt.Println("Register -- grpc -- imp")

	username := req.Username
	password := req.Password
	passwordConfirm := req.PasswordConfirm
	nickname := req.Nickname

	service := service.UserRegisterService{
		Nickname:        nickname,
		UserName:        username,
		Password:        password,
		PasswordConfirm: passwordConfirm,
	}

	serviceRes := service.Register()

	fmt.Println("注册结果 ",serviceRes)

	res.Code = int32(serviceRes.Code)
	res.Message = serviceRes.Msg

	return nil
}

func (a *Account) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	fmt.Println("login -- grpc -- imp")

	username := req.Username
	password := req.Password

	service := service.UserLoginService{
		UserName:        username,
		Password:        password,
	}

	serviceRes := service.LoginFromGrpc()

	res.Code = int32(serviceRes.Code)
	res.Message = serviceRes.Msg

	return nil
}
func (a *Account) Me(ctx context.Context, req *proto.MeRequest, res *proto.MeResponse) error {
	fmt.Println("Me -- grpc -- imp")

	return nil
}
func (a *Account) Logout(ctx context.Context, req *proto.LogoutRequest, res *proto.LogoutResponse) error {
	fmt.Println("Logout -- grpc -- imp")


	return nil
}