package models

import (
	"database/sql"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var sqlDB *sql.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) error {

	openDB, err := sql.Open("mysql",connString)

	// Error
	if err != nil {
		return errors.Wrap(err,"mysql connect failed")
	}

	sqlDB = openDB

	return nil
}

func InstanceDB() *sql.DB {
	return sqlDB
}
