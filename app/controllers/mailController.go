package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezor-dev/nezora/backend/app/models"
	"github.com/nezor-dev/nezora/backend/pkg/utils"
	"github.com/nezor-dev/nezora/backend/platform/database"
)

func GetMail(c *fiber.Ctx) error {
	id := c.Params("id")
	var mail models.Mail

	result := database.DB.Find(&mail, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "mail with the given ID is not found",
			"data":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  mail,
	})
}

func GetAllMail(c *fiber.Ctx) error {
	var mail []models.Mail

	database.DB.Find(&mail)
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(mail),
		"data":  mail,
	})
}

func CreateMail(c *fiber.Ctx) error {
	// initialise new model struc mail
	mail := new(models.Mail)

	// parse json to mail model
	err := c.BodyParser(mail)
	if err != nil {
		// return error 400 and error message
		c.Status(400).JSON(&fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return err
	}

	// Create a new validator for a mail model.
	validate := utils.NewValidator()
	// Validate mail fields.
	if err := validate.Struct(mail); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create mail
	database.DB.Create(&mail)
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
		"data":    mail,
	})
	return nil
}
func DeleteMail(c *fiber.Ctx) error {
	//initialize model mail from struct
	mail := new(models.Mail)

	// parse json data to struct
	if err := c.BodyParser(mail); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// set Mail id
	id := mail.ID

	// Delete mail
	result := database.DB.Delete(&mail, id)

	if result.RowsAffected == 0 {
		// Return status 500 and error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "mail entry not found",
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
func UpdateMail(c *fiber.Ctx) error {
	//initialize model mail from struct
	mail := new(models.Mail)

	// parse json data to struct
	if err := c.BodyParser(mail); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// set Mail id
	id := mail.ID

	// Create a new validator for a mail model.
	validate := utils.NewValidator()

	// Validate mail fields.
	if err := validate.Struct(mail); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update mail by given ID.
	result := database.DB.Where("id = ?", id).Updates(&mail)
	if result.RowsAffected == 0 {
		// Return status 500 if no rows affected
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "mail entry not found",
		})
	}
	// Return status 201.
	return c.Status(fiber.StatusCreated).JSON(mail)
}
