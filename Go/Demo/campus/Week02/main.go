package main

import (
	"Week02/conf"
	"Week02/server"
)

func main() {
	// 配置初始化
	err := conf.Init()
	if err != nil {
		panic(err)
		return
	}

	// 装载路由
	r := server.NewRouter()
	r.Run(":8080")
}
