package main

import (
	"go_wallpaper/conf"
	"go_wallpaper/server"
	"google.golang.org/grpc"
)

func main() {

	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":8080")

	//1、Dail连接
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

}
