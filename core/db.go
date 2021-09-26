package core

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"myblog/lib/config"
	"gorm.io/gorm/logger"
	"myblog/lib/helper"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	var (
		host     = config.Env("database.mysql.host")
		port     = config.Env("database.mysql.port")
		database = config.Env("database.mysql.database")
		username = config.Env("database.mysql.username")
		password = config.Env("database.mysql.password")
		charset  = config.Env("database.mysql.charset")
	)
	dsn :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")
	gormConfig := mysql.New(mysql.Config{
		DSN:dsn,
	})

	var level logger.LogLevel

	if helper.ToBool(config.Env("app.debug")){
		level = logger.Warn	// 读取不到数据也会报错
	}else{
		level = logger.Error// 发生错误就报错
	}
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: logger.Default.LogMode(level),
	}) // 数据库连接池

	helper.LogError(err)
	return DB
}