package main

import (
	"AccountServer/infra/base"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main() {
	// 获取配置文件
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)

}
