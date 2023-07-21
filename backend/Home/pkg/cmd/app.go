package cmd

import (
	"Home/pkg/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Run() {
	app := fiber.New(fiber.Config{
		AppName: "Home Service",
	})

	route.SetupRoutes(app)

	err := app.Listen(":3001")
	if err != nil {
		log.Fatal(err)
	}
}
