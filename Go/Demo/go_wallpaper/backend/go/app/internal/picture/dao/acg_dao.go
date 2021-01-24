package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_wallpaper/backend/go/internal/picture/model"
)

type AcgDao struct {
	db *gorm.DB
}

func NewAcgDao(db *gorm.DB) *AcgDao {
	//fmt.Println("new picture dao ", db)
	d := &AcgDao{
		db: db,
	}
	d.CreateAcgTable()
	return d
}

/*查询*/
func (d *AcgDao) GetPictures(page int, pageSize int) ([]model.Picture, error) {

	Pics := []model.Picture{}

	err := d.db.Scopes(Paginate(page, pageSize)).Find(&Pics).Error
	if err != nil {
		return nil, err
	}

	return Pics, nil
}

/*添加*/
func (d *AcgDao) AddAcg(data *model.Acg) error {
	// 以image url 为key 去重
	err := d.db.Where(model.Acg{ImageUrl: data.ImageUrl}).FirstOrCreate(&data).Error

	if err != nil {
		fmt.Println("新增照片数据错误 ", err)
		return err
	}

	return nil
}

func (d *AcgDao) Exist(picID string) (bool, error) {

	var pic model.Acg

	err := d.db.Where("picture_id = ?", picID).First(&pic).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *AcgDao) CreateAcgTable() {

	if !d.db.HasTable(&model.Acg{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.Acg{}).Error; err != nil {
			fmt.Println("建表错误,err", err)
		}
	}

}
