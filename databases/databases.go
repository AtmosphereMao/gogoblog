package databases

import (
	"gorm.io/gorm"
	"myblog/app/models/article"
	"myblog/app/models/category"
	"myblog/app/models/user"
	"myblog/core"
	"myblog/lib/config"
	"myblog/lib/helper"
	"time"
)

var DB *gorm.DB

func InitDB(){
	db := core.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(helper.ToInt(config.Env("database.mysql.max_open_connections")))
	sqlDB.SetMaxIdleConns(helper.ToInt(config.Env("database.mysql.max_idle_connections")))
	sqlDB.SetConnMaxLifetime(time.Duration(helper.ToInt(config.Env("database.mysql.max_life_seconds"))) * time.Second)

	migration(db)
}

func migration(db *gorm.DB){
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{})
}