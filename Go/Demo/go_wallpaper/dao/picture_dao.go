package dao

import (
	"fmt"
	"go_wallpaper/model"
	"gorm.io/gorm"
)

type PictureDao struct {
	db *gorm.DB
}

func NewPictureDao(db *gorm.DB) *PictureDao {
	return &PictureDao{
		db: db,
	}
}

/*查询*/
func (d *PictureDao) GetPictures(page int, pageSize int) (error, []model.Picture) {

	return nil, nil
}

/*添加*/
func (d *PictureDao) AddPicture(data *model.Picture) error {

	d.CreateTable()
	err := d.db.Create(&data).Error

	if err != nil {
		fmt.Println("新增照片数据错误,err", err)

	}
	fmt.Println("插入照片成功")

	return err
}

func (d *PictureDao) CreateTable() {

	if !d.db.HasTable(&model.Picture{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Photo{}).Error; err != nil {
			fmt.Println("建表错误,err", err)
		}
	}

}
