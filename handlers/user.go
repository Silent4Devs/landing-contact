package handlers

import (
	"fiber-boilerplate/databases"
	"fiber-boilerplate/models"
	"fiber-boilerplate/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := databases.Database.Find(&users).Error; err != nil {
		utils.Logger.Error("Failed to retrieve users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve users",
		})
	}
	utils.Logger.Info("Retrieved users successfully", zap.Int("count", len(users)))
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := databases.Database.First(&user, id).Error; err != nil {
		utils.Logger.Warn("User not found", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	utils.Logger.Info("Retrieved user successfully", zap.String("id", id))

	return c.Status(fiber.StatusOK).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)

	// Parse request body
	if err := c.BodyParser(user); err != nil {
		utils.Logger.Error("Invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Save to database
	if err := databases.Database.Create(&user).Error; err != nil {
		utils.Logger.Error("Failed to create user", zap.Any("user", user), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create user",
		})
	}

	utils.Logger.Info("User created successfully", zap.Any("user", user))
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// Check if the user exists
	if err := databases.Database.First(&user, id).Error; err != nil {
		utils.Logger.Warn("User not found for update", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Parse request body
	updateData := new(models.User)
	if err := c.BodyParser(updateData); err != nil {
		utils.Logger.Error("Invalid request body for update", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update the user in the database
	if err := databases.Database.Model(&user).Updates(updateData).Error; err != nil {
		utils.Logger.Error("Failed to update user", zap.String("id", id), zap.Any("updateData", updateData), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update user",
		})
	}

	utils.Logger.Info("User updated successfully", zap.String("id", id), zap.Any("updateData", updateData))
	return c.Status(fiber.StatusOK).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Delete the user
	if err := databases.Database.Delete(&models.User{}, id).Error; err != nil {
		utils.Logger.Warn("Failed to delete user", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	utils.Logger.Info("User deleted successfully", zap.String("id", id))
	return c.SendStatus(fiber.StatusOK)
}
