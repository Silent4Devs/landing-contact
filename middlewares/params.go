package middlewares

import (
	// Your custom logger
	"fiber-boilerplate/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

// Params sets up middlewares for the application, including logging with zap.
func Params(a *fiber.App) {
	// Use custom zap logger middleware to log incoming requests
	a.Use(func(c *fiber.Ctx) error {
		// Log the incoming request with zap
		utils.Logger.Info("Incoming request", zap.String("method", c.Method()), zap.String("path", c.Path()))

		// Continue to the next middleware/handler
		return c.Next()
	})

	// Use CORS middleware with credentials allowed
	a.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Use recovery middleware to handle panics
	a.Use(recover.New())
}
