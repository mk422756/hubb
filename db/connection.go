package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")

	dbTemp, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	dbTemp.AutoMigrate(&User{})
	dbTemp.AutoMigrate(&Page{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db = dbTemp
}

func Close() {
	db.Close()
}

func GetDB() *gorm.DB {
	return db
}
