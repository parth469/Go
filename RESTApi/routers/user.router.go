package routes

import (
	"github.com/gofiber/fiber/v3"
	controller "github.com/parth469/controller"
)

func userRoutes(App fiber.Router) {
	App.Get("/users", controller.GetAllUser)
	App.Get("/users/:id", controller.GetUser)

	App.Post("/users", controller.CreateUser)

	App.Patch("/users/:id", controller.UpdateUser)

	App.Delete("/users/:id", controller.DeleteUser)
}
