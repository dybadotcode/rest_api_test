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
func ConnectDB(dialect string, params string) {
	db, err := gorm.Open(dialect, params)
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Rss{})
	DB = db
	//DB.LogMode(true)
}
