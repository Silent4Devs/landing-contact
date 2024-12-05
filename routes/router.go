package routes

import (
	"fiber-boilerplate/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlers.Welcome)
	// Middleware
	api := app.Group("/api")

	//Index endpoint
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
	api.Get("/authenticated", handlers.AuthenticatedUser)
	api.Post("/logout", handlers.Logout)
	api.Post("/forgot", handlers.Forgot)
	api.Post("/reset", handlers.ResetPassword)
	api.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	//Users
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUser)
	api.Post("/users", handlers.AddUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.RemoveUser)

}
