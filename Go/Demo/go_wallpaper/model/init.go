package model

import (
	"fmt"
	"go_wallpaper/util"
	"time"

	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var db *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {

	fmt.Println("开始连接数据库...")
	fmt.Println(connString)

	opendb, err := gorm.Open("mysql", connString)
	db = opendb
	db.LogMode(true)

	fmt.Println("数据库初始化完成...")

	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(time.Hour)

	//migration()
}
