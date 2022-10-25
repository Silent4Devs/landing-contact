package handlers

import (
	"robot-monitoreo/databases"
	"robot-monitoreo/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	databases.Database.Create(&user)

	return c.Status(201).JSON(user)
}
