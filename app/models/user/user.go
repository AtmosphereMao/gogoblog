package user

import (
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
	if err := core.DB.Create(&user).Error;err != nil{
		helper.LogError(err)
		return err
	}
	return nil
}