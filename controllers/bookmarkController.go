package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/initializers"
	"github.com/nezor-dev/nezora/backend/models"
)

func GetBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	var bookmark models.Bookmark

	result := initializers.DB.Find(&bookmark, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(bookmark)
}

func GetAllBookmarks(c *fiber.Ctx) error {
	var bookmarks []models.Bookmark

	initializers.DB.Find(&bookmarks)
	return c.Status(200).JSON(bookmarks)
}

func CreateBookmark(c *fiber.Ctx) error {
	newBookmark := new(models.Bookmark)

	err := c.BodyParser(newBookmark)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	initializers.DB.Create(&newBookmark)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    newBookmark,
	})
	return nil
}
func DeleteBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	var bookmark models.Bookmark

	result := initializers.DB.Delete(&bookmark, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
func UpdateBookmark(c *fiber.Ctx) error {
	bookmark := new(models.Bookmark)
	id := c.Params("id")

	if err := c.BodyParser(bookmark); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	initializers.DB.Where("id = ?", id).Updates(&bookmark)
	return c.Status(200).JSON(bookmark)
}
