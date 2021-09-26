package article

import (
	"myblog/models"
	"myblog/models/user"
)

type Article struct{
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;"valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User

	CategoryID uint64 `gorm:"not null;default:4;index"`
}