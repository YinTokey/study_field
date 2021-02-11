package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"unsplash_server/cache"
	"unsplash_server/model"
	"unsplash_server/util"
)

// Init 初始化配置项
func Init() {
	//  配置文件加载到 viper中
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("unsplash_server") // TODO: 后面数据调整后，要改为 unsplash_server

	// 从本地读取环境变量
	//godotenv.Load()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("配置文件加载失败", err)
	}

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	dbUsername := viper.Get("db.username")
	dbPassword := viper.Get("db.password")
	dbName := viper.Get("db.name")
	dbAddr := viper.Get("db.addr")
	dbConfig := viper.Get("db.config")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", dbUsername, dbPassword, dbAddr, dbName, dbConfig)

	fmt.Println(dsn)

	// 连接数据库
	model.Database(dsn)
	cache.Redis()
}
