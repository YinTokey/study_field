package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_wallpaper/internal/picture/model"
)

type PicsumDao struct {
	db *gorm.DB
}

func NewPicsumDao(db *gorm.DB) *PicsumDao {
	//fmt.Println("new picture dao ", db)
	d := &PicsumDao{
		db: db,
	}
	d.CreatePicsumTable()
	return d
}

/*查询*/
func (d *PicsumDao) GetPictures(page int, pageSize int) ([]model.Picsum, error) {

	Pics := []model.Picsum{}

	err := d.db.Scopes(Paginate(page, pageSize)).Find(&Pics).Error
	if err != nil {
		return nil, err
	}

	return Pics, nil
}

/*添加*/
func (d *PicsumDao) AddPicsum(data *model.Picsum) error {
	// 以image url 为key 去重
	err := d.db.Where(model.Picsum{ImageUrl: data.ImageUrl}).FirstOrCreate(&data).Error

	if err != nil {
		fmt.Println("新增照片数据错误 ", err)
		return err
	}

	return nil
}

func (d *PicsumDao) Exist(picID string) (bool, error) {

	var pic model.Picsum

	err := d.db.Where("picture_id = ?", picID).First(&pic).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *PicsumDao) CreatePicsumTable() {

	if !d.db.HasTable(&model.Picsum{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.Acg{}).Error; err != nil {
			fmt.Println("建表错误,err", err)
		}
	}

}
