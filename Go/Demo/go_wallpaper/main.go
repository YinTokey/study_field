package main

import (
	"go_wallpaper/conf"
	"go_wallpaper/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":8080")
}
