package models

import (
	"time"

	"github.com/tony-skywalks/my-web/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName string    `json:firstname`
	LastName  string    `json:lastname`
	Username  string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    `json:password`
	LastLogin time.Time `json:lastlogin`
}

type Response struct {
	Staus string      `json:status`
	Error string      `json:error`
	Data  interface{} `josn:data`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (usr *User) CreateUser() (*User, error) {
	usr.Password, _ = usr.hashPassword(usr.Password)
	res := db.Create(&usr)
	return usr, res.Error
}

func (usr *User) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (usr User) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (usr *User) Autheticate() *User {
	var res User
	if err := db.Where("username = ?", usr.Username).First(&res).Error; err != nil {
		return &User{}
	}
	if !usr.CheckPassword(usr.Password, res.Password) {
		return &User{}
	}
	return &res
}
