package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nezor-dev/nezora/backend/app/models"
	"github.com/nezor-dev/nezora/backend/pkg/utils"
	"github.com/nezor-dev/nezora/backend/platform/database"
)

func GetBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	var bookmark models.Bookmark

	result := database.DB.Find(&bookmark, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "bookmark with the given ID is not found",
			"data":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  bookmark,
	})
}

func GetAllBookmarks(c *fiber.Ctx) error {
	var bookmarks []models.Bookmark
	// Get all bookmarks
	database.DB.Find(&bookmarks)
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(bookmarks),
		"data":  bookmarks,
	})
}

func CreateBookmark(c *fiber.Ctx) error {
	// initialise new model struc bookmark
	newBookmark := new(models.Bookmark)

	// parse json to bookmark model
	err := c.BodyParser(newBookmark)
	if err != nil {
		// return error 400 and error message
		c.Status(400).JSON(&fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return err
	}

	newBookmark.ID = uuid.New()

	// Create a new validator for a Bookmark model.
	validate := utils.NewValidator()
	// Validate bookmark fields.
	if err := validate.Struct(newBookmark); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create bookmark
	database.DB.Create(&newBookmark)
	if err != nil {
		// return status 500 and error message
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return err
	}

	// return status 200 OK
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    newBookmark,
	})
	return nil
}
func DeleteBookmark(c *fiber.Ctx) error {
	// Create new Book struct
	bookmark := &models.Bookmark{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(bookmark); err != nil {
		// Return status 400 and error message.
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Validate only one bookmark field ID.
	if err := validate.StructPartial(bookmark, "id"); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "bookmark with id not found",
		})
	}

	// Delete mail
	database.DB.Delete(&bookmark, bookmark.ID)

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
func UpdateBookmark(c *fiber.Ctx) error {

	//initialize model bookmark from struct
	bookmark := new(models.Bookmark)

	// parse json data to struct
	if err := c.BodyParser(bookmark); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// set bookmark id
	id := bookmark.ID

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Validate bookmark fields.
	if err := validate.Struct(bookmark); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update book by given ID.
	result := database.DB.Where("id = ?", id).Updates(&bookmark)
	if result.RowsAffected == 0 {
		// Return status 500 if no rows affected
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "bookmark with the given ID is not found",
		})
	}

	// Return status 201.
	return c.Status(fiber.StatusCreated).JSON(bookmark)
}
