package main

import (
	"fiber-boilerplate/config"
	"fiber-boilerplate/databases"
	"fiber-boilerplate/middlewares"
	"fiber-boilerplate/pkg/utils"
	"fiber-boilerplate/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

func main() {
	// Initialize the logger
	utils.InitLogger()
	defer utils.Logger.Sync() // Ensure logger flushes any pending logs on exit

	// Log application start
	utils.Logger.Info("Starting application", zap.String("timestamp", carbon.Now().ToDateTimeString()))

	carbon.Now().ToDateTimeString()
	// Define Fiber config.
	config := config.FiberConfig()

	//start fiber
	app := fiber.New(config)

	// Middlewares.
	utils.Logger.Info("Applying middlewares...")
	middlewares.Params(app)

	//connect database
	utils.Logger.Info("Connecting to the database...")
	databases.Connect()

	utils.Logger.Info("Setting up routes...")
	routes.SetupRoutes(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.Logger.Info("Starting server in development mode...")
		utils.StartServer(app)
	} else {
		utils.Logger.Info("Starting server with graceful shutdown...")
		utils.StartServerWithGracefulShutdown(app)
	}

}
