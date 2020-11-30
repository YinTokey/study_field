package models

import (
	"database/sql"
	"fmt"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var db *gorm.DB

var sqlDB *sql.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	fmt.Println(connString)

	fmt.Println("开始连接数据库...")

	openDB, err := sql.Open("mysql",connString)

	//opendb, err := gorm.Open("mysql", connString)
	//
	//db = opendb

	// Error
	if err != nil {
		fmt.Println("连接数据库不成功 ")
		fmt.Println(err)

		//util.Log().Panic("连接数据库不成功", err)
	}

	sqlDB = openDB

}

func InstanceDB() *sql.DB {
	fmt.Println("获取单例")
	fmt.Println(db)
	return sqlDB
}
