package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/controllers"
)

func AddDocumentRoutes(app *fiber.App) {
	app.Get("/documents", controllers.GetAllDocuments)
	app.Get("/documents/:id", controllers.GetDocument)
	app.Post("/documents", controllers.CreateDocument)
	app.Put("/documents/:id", controllers.UpdateDocument)
	app.Delete("/documents/:id", controllers.DeleteDocument)
}
