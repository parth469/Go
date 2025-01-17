package main

import (
	"fmt"

	server "github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/log"
)

func main() {
	app := server.New(server.Config{AppName: "Learning GO Fiber"})

	app.Post("/post", func(c *server.Ctx) error {
		payload := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		return c.JSON(payload)
	})

	app.Listen(":3000")
}

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func getName(c *server.Ctx) error {
	params := c.AllParams()
	p := new(Person)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	fmt.Println(params, p.Name)
	return c.SendString("You get Nmae")
}
func getDetail(c *server.Ctx) error {
	logger.Info("Hello, World rebuild !")
	return c.SendString("Hello, World!")
}
