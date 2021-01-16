package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_wallpaper/internal/comment/model"
)

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	fmt.Println("new picture dao ", db)
	d := &CommentDao{
		db: db,
	}
	d.CreatePicTable()
	return d
}

func (d *CommentDao) CreatePicTable() {
	//  subject 表
	if !d.db.HasTable(&model.CommentSubject{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建 subject 表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.CommentSubject{}).Error; err != nil {
			fmt.Println("建 subject 表错误,err", err)
		}
	}

	// index 表
	if !d.db.HasTable(&model.CommentIndex{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建 CommentIndex 表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.CommentIndex{}).Error; err != nil {
			fmt.Println("建 CommentIndex 表错误,err", err)
		}
	}

	// content 表
	if !d.db.HasTable(&model.CommentContent{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建 CommentContent 表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.CommentContent{}).Error; err != nil {
			fmt.Println("建 CommentContent 表错误,err", err)
		}
	}
}

//func (d *CommentDao) CreatePicTable() {
//
//}
