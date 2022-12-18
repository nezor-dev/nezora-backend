package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/app/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/bookmark", controllers.CreateBookmark) // create a new bookmark

	// Routes for PUT method:
	route.Put("/bookmark", controllers.UpdateBookmark) // update one bookmark by ID

	// Routes for DELETE method:
	route.Delete("/bookmark", controllers.DeleteBookmark) // delete one bookmark by ID

	// Routes for POST method:
	route.Post("/mail", controllers.CreateMail) // create a new mail

	// Routes for PUT method:
	route.Put("/mail", controllers.UpdateMail) // update one mail by ID

	// Routes for DELETE method:
	route.Delete("/mail", controllers.DeleteMail) // delete one mail by ID
}
