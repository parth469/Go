package main

import (
	fiber "github.com/gofiber/fiber/v3"
	config "github.com/parth469/config"
	"github.com/parth469/controller"
	database "github.com/parth469/database"
	routes "github.com/parth469/routers"
	logger "github.com/parth469/utils"
)

func main() {

	logger.InitLogger()
	defer logger.SyncLogger()

	cnf, err := config.LoadConfig()
	if err != nil {
		logger.Message("Failed to load environment variables", logger.Error)
		logger.SyncLogger()
		return
	}
	logger.Message("Environment variables loaded successfully", logger.Info)

	DBConnectError := database.ConnectDatabase(cnf.Database)

	if DBConnectError != nil {
		logger.Message("Failed to connect to the database: "+DBConnectError.Error(), logger.Error)
		logger.SyncLogger()
		return
	}
	logger.Message("Database connected successfully", logger.Info)
	controller.CreateValidator()

	app := fiber.New(fiber.Config{AppName: cnf.AppName})

	routes.CustomRoutes(app)

	if err := app.Listen(":" + cnf.Port); err != nil {
		logger.Message("Failed to start server: "+err.Error(), logger.Error)
	}
}
