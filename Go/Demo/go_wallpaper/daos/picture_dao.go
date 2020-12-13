package daos

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_wallpaper/model"
)

type PictureDao struct {
	db *gorm.DB
}

func NewPictureDao(db *gorm.DB) *PictureDao {
	d := &PictureDao{
		db: db,
	}
	d.CreatePicTable()
	return d
}

/*查询*/
func (d *PictureDao) GetPictures(page int, pageSize int) (error, []model.Picture) {

	return nil, nil
}

/*添加*/
func (d *PictureDao) AddPicture(data model.Picture) error {

	err := d.db.Where(model.Picture{PictureId: data.PictureId}).FirstOrCreate(&data).Error

	if err != nil {
		fmt.Println("新增照片数据错误 ", err)
		return err
	}

	return nil
}

func (d *PictureDao) Exist(picID string) (bool, error) {

	var pic model.Picture

	err := d.db.Where("picture_id = ?", picID).First(&pic).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *PictureDao) CreatePicTable() {

	if !d.db.HasTable(&model.Picture{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.Picture{}).Error; err != nil {
			fmt.Println("建表错误,err", err)
		}
	}

}
