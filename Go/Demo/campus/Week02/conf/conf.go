package conf

import (
	"Week02/models"
	"os"
)

// Init 初始化配置项
func Init() {
	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
}
