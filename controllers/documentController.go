package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/initializers"
	"github.com/nezor-dev/nezora/backend/models"
)

func GetDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	var Document models.Document

	result := initializers.DB.Find(&Document, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(Document)
}

func GetAllDocuments(c *fiber.Ctx) error {
	var Documents []models.Document

	initializers.DB.Find(&Documents)
	return c.Status(200).JSON(Documents)
}

func CreateDocument(c *fiber.Ctx) error {
	newDocument := new(models.Document)

	err := c.BodyParser(newDocument)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	initializers.DB.Create(&newDocument)

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
		"data":    newDocument,
	})
	return nil
}
func DeleteDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	var document models.Document

	result := initializers.DB.Delete(&document, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
func UpdateDocument(c *fiber.Ctx) error {
	Document := new(models.Document)
	id := c.Params("id")

	if err := c.BodyParser(Document); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	initializers.DB.Where("id = ?", id).Updates(&Document)
	return c.Status(200).JSON(Document)
}
