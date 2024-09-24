package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/parth469/models"
	"github.com/parth469/service"
	logger "github.com/parth469/utils"
)

func GetAllUser(c fiber.Ctx) error {
	detail, err := service.AllUserFetch()
	if err != nil {
		logger.Message(err.Error(), logger.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	messess := "You call Get All User Function"
	logger.Message(messess, logger.Info)
	return c.JSON(detail)
}

func CreateUser(c fiber.Ctx) error {

	body := c.Body()

	var user models.UserCreateInput

	if len(body) == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": "Body Is Missing"})
	}

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if err := validate.Struct(user); err != nil {
		formattedErrors := formatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": formattedErrors})
	}
	service.CreateUser(user)
	return nil
}

func UpdateUser(c fiber.Ctx) error {
	return nil

}

func GetUser(c fiber.Ctx) error {
	return nil

}

func DeleteUser(c fiber.Ctx) error {
	return nil

}
