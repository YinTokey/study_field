package base

import (
	"AccountServer/infra"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
)

var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

type DbxDataBaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDataBaseStarter) Setup(ctx infra.StarterContext) {
	// 获取初始配置
	conf := ctx.Props()

	settings := dbx.Settings{}
	// 把配置解析到 settings 结构体中
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Info("mysql.conn url:", settings.ShortDataSourceName())
	database, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(database.Ping())

}
