package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
	"unsplash_server/global"
	"unsplash_server/pkg/logger"
)

type AppSettingS struct {
	defaultContextTimeout time.Duration
	logSavePath           string
	logFileName           string
	logFileExt            string
}

// Init 初始化配置项
func Init() {
	//  配置文件加载到 viper中
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("unsplash_server") // TODO: 后面数据调整后，要改为 unsplash_server

	// 从本地读取环境变量
	//godotenv.Load()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("配置文件加载失败", err)
	}


	dbUsername := viper.Get("db.username")
	dbPassword := viper.Get("db.password")
	dbName := viper.Get("db.name")
	dbAddr := viper.Get("db.addr")
	dbConfig := viper.Get("db.config")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", dbUsername, dbPassword, dbAddr, dbName, dbConfig)

	fmt.Println(dsn)

	_ = setupLogger()

	//// 连接数据库
	//model.Database(dsn)
	//cache.Redis()
}

func setupLogger() error {

	logSavePath := viper.Get("app.logSavePath").(string)
	logFileName := viper.Get("app.logFileName").(string)
	logFileExt := viper.Get("app.logFileExt").(string)

	fileName := logSavePath + "/" + logFileName + logFileExt

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}