package base

import (
	"AccountServer/infra"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
)

//dbx 数据库实例
var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

//dbx数据库starter，并且设置为全局
type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDatabaseStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	//数据库配置
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Info("mysql.conn url:", settings.ShortDataSourceName())
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx.Ping())
	database = dbx
}
