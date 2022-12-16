package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	var MARIADB_DATABASE = os.Getenv("MARIADB_DATABASE")
	var MARIADB_USER = os.Getenv("MARIADB_USER")
	var MARIADB_PASSWORD = os.Getenv("MARIADB_PASSWORD")
	var GO_DB_HOST = os.Getenv("GO_DB_HOST")

	dsn := MARIADB_USER + ":" + MARIADB_PASSWORD + "@tcp(" + GO_DB_HOST + ")/" + MARIADB_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
}
