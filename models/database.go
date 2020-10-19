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
<<<<<<< HEAD
func ConnectDB(dialect string, params string) {
	db, err := gorm.Open(dialect, params)
=======
func ConnectDB() {
	db, err := gorm.Open("postgres", "host= port= user= dbname= password= sslmode=disable")
>>>>>>> 81f14c17c98ebab3d7f31e7daa6056ead0e538fe
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Rss{})
	DB = db
	//DB.LogMode(true)
}
