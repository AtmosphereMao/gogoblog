package user

import (
	"encoding/base64"
	"github.com/elithrar/simple-scrypt"
	"myblog/app/models"
	"myblog/core"
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
	dk, _ := scrypt.GenerateFromPassword([]byte(user.Password),scrypt.DefaultParams)
	user.Password = base64.StdEncoding.EncodeToString(dk)
	if err := core.DB.Create(&user).Error;err != nil{
		helper.LogError(err)
		return err
	}
	return nil
}

func (user *User)GetByEmail()(User, error){
	var _user User
	if err := core.DB.Where("email = ?", user.Email).First(&_user).Error; err != nil{
		return _user, err
	}
	return _user, nil
}

func (user *User)EditPassword()(error){
	dk, _ := scrypt.GenerateFromPassword([]byte(user.Password),scrypt.DefaultParams)
	user.Password = base64.StdEncoding.EncodeToString(dk)
	if err := core.DB.Save(&user).Error;err != nil{
		helper.LogError(err)
		return err
	}
	return nil
}