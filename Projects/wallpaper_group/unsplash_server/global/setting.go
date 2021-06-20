package global

import (
	"unsplash_server/pkg/logger"
	"unsplash_server/pkg/setting"
)

// 存放全局单例

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettingS
	Logger          *logger.Logger
)
