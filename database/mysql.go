package database

import (
	"fmt"
	"log"
	"mini-project/models"
	"mini-project/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		util.Cfg.DB_USERNAME,
		util.Cfg.DB_PASSWORD,
		util.Cfg.DB_HOST,
		util.Cfg.DB_PORT,
		util.Cfg.DB_NAME,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("fail to connect database ", err)
	}

	MigrateDB(DB)
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Blog{},
		&models.Like{},
		&models.Comment{},
	)
}
