package repository

import (
	"errors"
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
