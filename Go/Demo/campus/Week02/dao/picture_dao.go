package dao

import (
	"Week02/models"
	"database/sql"
	"fmt"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type PictureDao struct {
	db *sql.DB
}

func NewPictureDao(db *sql.DB) *PictureDao {
	return &PictureDao{
		db: db,
	}
}

func (d *PictureDao) AddPicture(data *models.Picture) error {
	//err := d.CreateTable()
	//if err != nil {
	//	return err
	//}
	//fmt.Println(data)
	//createDb := d.db.Create(&data)
	//return createDb.Error
	//
	return nil
}

func (d *PictureDao) GetAll() (models.Picture, error) {

	var result models.Picture

	var author string

	err := d.db.QueryRow("select author from pictures where id = ?", 1).Scan(&author)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("no user with id %d\n", 1)
	case err != nil:
		fmt.Println("query error: %v\n", err)
	default:
		fmt.Println("author is %q", author)
	}

	return result, err

}

func (d *PictureDao) CreateTable() error {

	//if !d.db.HasTable(&models.Picture{}) {
	//	//直接通过 db.CreateTable 就可以创建表了，非常方便，
	//	//还可以通过 db.Set 设置一些额外的表属性,
	//	fmt.Println("尝试建表 ")
	//
	//	err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.Picture{}).Error
	//	if err != nil {
	//		fmt.Println("建表错误,err", err)
	//
	//	}
	//
	//	return err
	//}
	return nil
}