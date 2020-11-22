package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
//type User struct {
//	gorm.Model
//	UserName       string
//	PasswordDigest string
//	Nickname       string
//	Status         string
//	Avatar         string `gorm:"size:1000"`
//}

type User struct {
	gorm.Model
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Fullname         string    `json:"fullname"`
	AvatarVersion    int       `json:"avatar_version"`
	RegistrationDate time.Time `json:"registration_date"`
	Avatars          struct {
		Tiny struct {
			HTTPS string `json:"https"`
		} `json:"tiny"`
		Small struct {
			HTTPS string `json:"https"`
		} `json:"small"`
		Large struct {
			HTTPS string `json:"https"`
		} `json:"large"`
		Default struct {
			HTTPS string `json:"https"`
		} `json:"default"`
	} `json:"avatars"`
	UserpicURL      string `json:"userpic_url"`
	UserpicHTTPSURL string `json:"userpic_https_url"`
	Usertype        int    `json:"usertype"`
	Active          int    `json:"active"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	About           string `json:"about"`
	City            string `json:"city"`
	State           string `json:"state"`
	Country         string `json:"country"`
	CoverURL        string `json:"cover_url"`
	UpgradeStatus   int    `json:"upgrade_status"`
	Affection       int    `json:"affection"`
	FollowersCount  int    `json:"followers_count"`
	Following       bool   `json:"following"`

	PasswordDigest   string

}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
