package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB
var err error

func InitMySql() {
	fmt.Println("init mysql")
	if db == nil {

		//db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/gitch?charset=utf8")
		//if err != nil {
		//	panic(err.Error())
		//}
		openConnector()

		CreateTableWithUser()
		CreateTableWithAlbum()
		CreateTableWithArticle()
	}
}

func openConnector() {
	Connector, err := mysql.NewConnector(&mysql.Config{
		User:   "root",
		Passwd: "12345678",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gitchat",
		AllowNativePasswords:true,
		Collation:"utf8_general_ci",
		ParseTime:true,
		Loc:time.Local,
	})
	if err != nil {
		panic(err)
	}
	db = sql.OpenDB(Connector)
	if err = db.Ping();err != nil{
		panic(err)
	}
}


func CreateTableWithUser() {
	sqlstr := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        createtime INT(10)
        );`

	ModifyDB(sqlstr)
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}


//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}
