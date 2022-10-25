package handlers

import (
	"robot-monitoreo/databases"
	"robot-monitoreo/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// user := new(models.User)
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(400).SendString("Password do not match!")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	databases.Database.Create(&user)

	return c.Status(200).JSON(user)
}
