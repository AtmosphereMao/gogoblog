package user

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"myblog/app/models"
	"myblog/core"
	"myblog/lib/config"
	"myblog/lib/helper"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (user *User)Create() (error){
	dk, _ := scrypt.Key([]byte(user.Password), []byte(helper.ToString(config.Env("PASSWORD_SALT"))), 1<<15, 8, 1, 32)
	user.Password = base64.StdEncoding.EncodeToString(dk)
	if err := core.DB.Create(&user).Error;err != nil{
		helper.LogError(err)
		return err
	}
	return nil
}