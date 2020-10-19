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
	db, err := gorm.Open("postgres", "host= port= user= dbname= password= sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Rss{})
	DB = db
	//DB.LogMode(true)
}
