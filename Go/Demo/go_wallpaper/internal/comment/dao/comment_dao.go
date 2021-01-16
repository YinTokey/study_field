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

func (d *CommentDao) QueryCommentSubject(id string) (*model.CommentSubject, error) {

	var subject *model.CommentSubject

	result := d.db.Where("obj_id = ?", id).First(&subject)

	return subject, result.Error
}

// 插入subject
func (d *CommentDao) AppendSubject(subject *model.CommentSubject) error {

	err := d.db.Where(model.CommentSubject{ObjId: subject.ObjId}).FirstOrCreate(&subject).Error

	if err != nil {
		fmt.Println("新增subject数据错误 ", err)
		return err
	}

	return nil

}

// 插入index
func (d *CommentDao) AppendIndex(index *model.CommentIndex) error {

	err := d.db.Where(model.CommentIndex{ObjId: index.ObjId}).FirstOrCreate(&index).Error

	if err != nil {
		fmt.Println("新增index数据错误 ", err)
		return err
	}

	return nil
}

// 插入content
func (d *CommentDao) AppendContent(content *model.CommentContent) error {

	err := d.db.Where(model.CommentContent{CommentId: content.CommentId}).FirstOrCreate(&content).Error

	if err != nil {
		fmt.Println("新增content数据错误 ", err)
		return err
	}
	fmt.Println("插入content 成功", content.Message)

	return nil
}

func (d *CommentDao) GetIndexs(obj_id string) ([]model.CommentIndex, error) {

	idxs := []model.CommentIndex{}

	err := d.db.Where("obj_id <> ?", obj_id).Find(&idxs).Error

	if err != nil {
		return nil, err
	}

	return idxs, nil
}

// 建表
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
