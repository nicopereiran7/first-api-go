package common

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/nicopereiran7/first-api-go/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1999_nico@/first_api_go_db?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando....")

	db.AutoMigrate(&models.User{})
}
