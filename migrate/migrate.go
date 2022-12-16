package main

import (
	"github.com/nezor-dev/nezora/backend/initializers"
	"github.com/nezor-dev/nezora/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Bookmark{})
	initializers.DB.AutoMigrate(&models.Document{})
}
