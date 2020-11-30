package dao

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)


type PictureDao struct {
	db *sql.DB
}

func NewPictureDao(db *sql.DB) *PictureDao {
	return &PictureDao{
		db: db,
	}
}

func (d *PictureDao) Query(id int) (string, error) {

	var author string

	err := d.db.QueryRow("select author from pictures where id = ?", id).Scan(&author)

	switch {
	case err == sql.ErrNoRows:
		return author, errors.Wrapf(err,"no user with id %d ", id)
	case err != nil:
		return author, errors.Wrapf(err,"uery error: %v ", err)
	default:
	}

	return author, err

}
