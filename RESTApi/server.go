package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	config "github.com/parth469/config"
	logger "github.com/parth469/log"
)

func main() {
	ctf, err := config.LoadConfig()
	logger.InitLogger()
	defer logger.SyncLogger()

	if err != nil {
		fmt.Println("Fail to Load Env error")
	}

	app := fiber.New(fiber.Config{AppName: ctf.AppName})

	app.Get("/", func(c fiber.Ctx) error {
		logger.Message("API CAll", logger.Error)
		return c.SendString("Optimize config file ")
	})
	app.Listen(":" + ctf.Port)
}
