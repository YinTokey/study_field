package main

import (
	"go_wallpaper/configs"
	"go_wallpaper/internal/server/http"
)

func main() {

	// 从配置文件读取配置
	configs.Init()

	// 装载路由
	r := http.NewRouter()
	r.Run(":8080")

}
