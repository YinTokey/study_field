package dao

import  (
	"Week02/models"
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
	//createDb := db.Create(&data)
	//return createDb.Error
	//err := createDb.Error {
	//
	//}
	return nil
}

func (d *PictureDao) GetAll() ([]models.Picture, error) {

	return nil, nil
}
