package cmd

import (
	"github.com/gofiber/fiber/v2"
	"helloWorld/pkg/route"
	"log"
)

func Run() {
	app := fiber.New(fiber.Config{
		AppName: "HelloWorld Server",
	})

	route.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
