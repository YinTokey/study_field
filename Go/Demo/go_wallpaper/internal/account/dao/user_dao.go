package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_wallpaper/internal/account/model"
	"golang.org/x/crypto/bcrypt"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	fmt.Println("new user dao ", db)
	d := &UserDao{
		db: db,
	}
	d.CreateUserTable()
	return d
}

func (d *UserDao) CreateUserTable() {

	if !d.db.HasTable(&model.User{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		if err := d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.User{}).Error; err != nil {
			fmt.Println("建表错误,err", err)
		}
	}

}

// GetUser 用ID获取用户
func (d *UserDao) GetUser(ID interface{}) (model.User, error) {
	var user model.User
	result := d.db.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (d *UserDao) SetPassword(user model.User, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), model.PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (d *UserDao) CheckPassword(user model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

/// 检查用户名是否存在
func (d *UserDao) CheckUserNameExist(username string) bool {
	var count = 0
	d.db.Model(&model.User{}).Where("user_name = ?", username).Count(&count)

	return count > 0
}

/// 检查昵称是否存在
func (d *UserDao) CheckNickNameExist(nickname string) bool {
	var count = 0
	d.db.Model(&model.User{}).Where("nickname = ?", nickname).Count(&count)

	return count > 0
}

/// 创建用户
func (d *UserDao) CreateUser(user *model.User) error {
	fmt.Println("CreateUser")
	if err := d.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
