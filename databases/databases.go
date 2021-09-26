package databases

import(
	"myblog/core"
	"myblog/lib/config"
	"myblog/lib/helper"
	"gorm.io/gorm"
	"myblog/models/article"
	"myblog/models/category"
	"myblog/models/user"
	"time"
)

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