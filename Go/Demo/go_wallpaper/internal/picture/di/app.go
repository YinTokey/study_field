package di

import (
	"go_wallpaper/internal/picture/server/http"
)

func InitApp() error {

	// 从配置文件读取配置
	//configs.Init()

	// grpc 监听
	//grpc.Serve()

	// 装载路由
	r := http.NewRouter()
	r.Run(":8080")

	return nil
}
