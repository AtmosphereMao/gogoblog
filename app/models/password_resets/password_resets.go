package password_resets

import (
	"myblog/core"
	"myblog/lib/helper"
	"time"
)

type PasswordResets struct {
	Email     string `gorm:"type:varchar(255)" valid:"email"`
	Token 	  string `gorm:"type:varchar(191)" valid:"token"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
}

func (pr *PasswordResets)Create()(error){
	if err := core.DB.Create(&pr).Error;err != nil{
		helper.LogError(err)
		return err
	}
	return nil
}

func (pr *PasswordResets)GetByEmail()(PasswordResets, error){
	var _pr PasswordResets
	if err := core.DB.Where("email = ?", pr.Email).First(&_pr).Error; err!=nil{
		return _pr, err
	}
	return _pr, nil
}

func (pr *PasswordResets)GetByToken()(PasswordResets, error){
	var _pr PasswordResets
	if err := core.DB.Where("token = ?", pr.Token).First(&_pr).Error; err!=nil{
		return _pr, err
	}
	return _pr, nil
}

func (pr *PasswordResets)DeleteByEmail()(error){
	if err := core.DB.Where("email = ?", pr.Email).Delete(&pr).Error; err != nil{
		helper.LogError(err)
		return err
	}
	return nil
}