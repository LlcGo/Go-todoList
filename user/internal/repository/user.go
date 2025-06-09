package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/logic"
)

type User struct {
	UserId         uint   `gorm:primarykey`
	UserName       string `gorm:unique`
	NickName       string
	PasswordDigest string
}

func (u *User) ShowUserInfo(req *logic.UserRequest) error {
	if exist := u.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

func (u *User) CheckUserExist(req *logic.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&u).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (u *User) UserCreate(req *logic.UserRequest) error {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exist")
	}
	user := User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassWord(req.Password)
	return DB.Create(&user).Error
}

func (u *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

func BuildUser(item User) *logic.UserModel {
	userModel := logic.UserModel{
		UserID:   uint32(item.UserId),
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}
