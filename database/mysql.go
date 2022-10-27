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
		util.GetConfig("DB_USERNAME"),
		util.GetConfig("DB_PASSWORD"),
		util.GetConfig("DB_HOST"),
		util.GetConfig("DB_PORT"),
		util.GetConfig("DB_NAME"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("fail to connect database ", err)
	}

	MigrateDB(DB)
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
}
