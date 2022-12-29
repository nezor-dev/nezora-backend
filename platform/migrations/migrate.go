package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nezor-dev/nezora/backend/app/models"
	"github.com/nezor-dev/nezora/backend/platform/database"
)

func init() {
	database.ConnectToDb()
}

func main() {
	database.DB.AutoMigrate(&models.Bookmark{})
	database.DB.AutoMigrate(&models.Mail{})
	database.DB.AutoMigrate(&models.User{})
}
