package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nezor-dev/nezora/backend/app/models"
	"github.com/nezor-dev/nezora/backend/pkg/configs"
	"github.com/nezor-dev/nezora/backend/pkg/routes"
	"github.com/nezor-dev/nezora/backend/pkg/utils"
	"github.com/nezor-dev/nezora/backend/platform/database"
)

func init() {
	database.ConnectToDb()

	database.DB.AutoMigrate(&models.Bookmark{})
	database.DB.AutoMigrate(&models.Mail{})
	database.DB.AutoMigrate(&models.User{})
}

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// configure cors
	app.Use(cors.New(configs.CorsConfig()))

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
