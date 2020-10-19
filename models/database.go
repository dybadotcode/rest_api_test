package models

import (
	// ORM framework
	"github.com/jinzhu/gorm"
	//HTTP framework
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ..
var DB *gorm.DB

//ConnectDB ...
func ConnectDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=demo password=q sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Rss{})
	DB = db
	DB.LogMode(true)
}
