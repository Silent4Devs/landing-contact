package handlers

import (
	"fmt"
	"robot-monitoreo/databases"
	"robot-monitoreo/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
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

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	var user models.User

	databases.Database.Where("email = ?", data["email"]).First(&user)

	fmt.Println(user.ID)
	if user.ID == 0 {
		return c.Status(503).SendString()
	}

	return c.Status(200).JSON(user)
}
