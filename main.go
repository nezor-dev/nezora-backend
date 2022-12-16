package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nezor-dev/nezora/backend/initializers"
	"github.com/nezor-dev/nezora/backend/models"
	"github.com/nezor-dev/nezora/backend/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	app := fiber.New()

	initializers.DB.AutoMigrate(&models.Bookmark{})
	initializers.DB.AutoMigrate(&models.Document{})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.LoadRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
