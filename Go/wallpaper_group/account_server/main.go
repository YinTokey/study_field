package main

import (
	"account_server/conf"
	"account_server/server"

)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// grpc 服务启动
	s := server.NewGrpcServer()
	go s.Start()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")

}
