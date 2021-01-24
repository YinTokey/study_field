package di

import (
	"fmt"
	"github.com/spf13/viper"
	"go_wallpaper/internal/picture/job/picsum"
	"go_wallpaper/pkg"

	"os"
)

func InitApp() error {

	// 从配置文件读取配置
	InitConfigs()

	// grpc 监听
	//grpc.Serve()

	// 装载路由
	//r := http.NewRouter()
	//r.Run(":8080")

	return nil
}

func InitConfigs() {

	// 从本地读取环境变量
	//godotenv.Load()

	// 设置日志级别
	pkg.BuildLogger(os.Getenv("LOG_LEVEL"))

	//  配置文件加载到 viper中
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("go_wallpaper") // 读取环境变量的前缀为 go_wallpaper

	if err := viper.ReadInConfig(); err != nil {
		pkg.Log().Panic("配置文件加载失败", err)
	}

	// 读取翻译文件
	//if err := LoadLocales("configs/locales/zh-cn.yaml"); err != nil {
	//	pkg.Log().Panic("翻译文件加载失败", err)
	//}

	dbUsername := viper.Get("db.username")
	dbPassword := viper.Get("db.password")
	dbName := viper.Get("db.name")
	dbAddr := viper.Get("db.addr")
	dbConfig := viper.Get("db.config")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", dbUsername, dbPassword, dbAddr, dbName, dbConfig)

	// 连接数据库
	pkg.Database(dsn)
	pkg.Redis()

	// 初始化发号器
	pkg.GuidGeneratorInit()

	j := picsum.NewPicsumJob()

	j.StartWork()

	//j := acg.NewRandomAcgJob()
	//
	//for i := 0; i < 10000; i++ {
	//	j.FetchJob_2()
	//}

	//for i := 1000; i < 4000; i++ {
	//	j.FetchJob_3(i)
	//}

}
