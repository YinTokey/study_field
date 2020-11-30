package conf

import (
	"Week02/models"
)

// Init 初始化配置项
func Init() error {

	// 连接数据库
	env := "root:12345678@tcp(127.0.0.1:3306)/week2_db?charset=utf8&parseTime=True&loc=Local"
	return models.Database(env)

}
