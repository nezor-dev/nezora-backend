package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/bookmarks", controllers.GetAllBookmarks) // get list of all bookmarks
	route.Get("/bookmark/:id", controllers.GetBookmark)  // get one bookmark by ID

	route.Get("/mails", controllers.GetAllMail) // get list of all mails
	route.Get("/mail/:id", controllers.GetMail) // get one mail by ID

	route.Post("/auth/login", controllers.Login)       // login
	route.Post("/auth/register", controllers.Register) // register
	route.Get("/user", controllers.User)               // user
	route.Get("/auth/logout", controllers.Logout)      //logout
}
