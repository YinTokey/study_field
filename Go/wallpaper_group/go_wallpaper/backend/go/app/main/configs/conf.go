package configs

// Init 初始化配置项
//func Init() {
// 从本地读取环境变量
//godotenv.Load()
//
//// 设置日志级别
//pkg.BuildLogger(os.Getenv("LOG_LEVEL"))
//
////  配置文件加载到 viper中
//viper.SetConfigName("config")
//viper.SetConfigType("yaml")
//viper.AddConfigPath("configs")
//viper.AutomaticEnv()
//viper.SetEnvPrefix("go_wallpaper") // 读取环境变量的前缀为 go_wallpaper
//
//// 监听配置文件变化
//watchConfig()
//
//if err := viper.ReadInConfig(); err != nil {
//	pkg.Log().Panic("配置文件加载失败", err)
//}
//
//// 读取翻译文件
////if err := LoadLocales("configs/locales/zh-cn.yaml"); err != nil {
////	pkg.Log().Panic("翻译文件加载失败", err)
////}
//
//dbUsername := viper.Get("db.username")
//dbPassword := viper.Get("db.password")
//dbName := viper.Get("db.name")
//dbAddr := viper.Get("db.addr")
//dbConfig := viper.Get("db.config")
//dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", dbUsername, dbPassword, dbAddr, dbName, dbConfig)
//
//// 连接数据库
//model.Database(dsn)
//dao.Redis()
//}

//func watchConfig() {

//viper.WatchConfig()
//viper.OnConfigChange(func(e fsnotify.Event) {
//	pkg.Log().Info("Config file changed: %s", e.Name)
//})

//}
