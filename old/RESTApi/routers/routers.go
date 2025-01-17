package routes

import (
	"github.com/gofiber/fiber/v3"
	logger "github.com/parth469/utils"
)

func CustomRoutes(App *fiber.App) {
	App.Get("/", func(c fiber.Ctx) error {
		message := " Welcome to Go REST API"
		logger.Message(message, logger.Info)
		return c.SendString(message)
	})

	V1Group := App.Group("/api/v1")

	userRoutes(V1Group)
}
