package dao

import  (
	"Week02/models"
	"fmt"
	"github.com/jinzhu/gorm"
	)


type PictureDao struct {
	db *gorm.DB
}

func NewPictureDao(db *gorm.DB) *PictureDao {
	return &PictureDao{
		db: db,
	}
}

func (d *PictureDao) AddPicture(data *models.Picture) error {
	err := d.CreateTable()
	if err != nil {
		return err
	}
	fmt.Println(data)
	createDb := d.db.Create(&data)
	return createDb.Error
}

func (d *PictureDao) GetAll() ([]models.Picture, error) {

	var result []models.Picture

	createDb := d.db.Find(&result)
	err := createDb.Error
	if err != nil {
		fmt.Println("查询所有photo数据错误,err", err)
		//panic(err)
	}
	fmt.Println(len(result))

	return result, err

}

func (d *PictureDao) CreateTable() error {

	if !d.db.HasTable(&models.Picture{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&models.Picture{}).Error
		if err != nil {
			fmt.Println("建表错误,err", err)

		}

		return err
	}

	return nil
}