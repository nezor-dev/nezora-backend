package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/controllers"
)

func AddBookmarkRoutes(app *fiber.App) {
	app.Get("/bookmarks", controllers.GetAllBookmarks)
	app.Get("/bookmarks/:id", controllers.GetBookmark)
	app.Post("/bookmarks", controllers.CreateBookmark)
	app.Put("/bookmarks/:id", controllers.UpdateBookmark)
	app.Delete("/bookmarks/:id", controllers.DeleteBookmark)
}
