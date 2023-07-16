package route

import (
	"github.com/gofiber/fiber/v2"
	"helloWorld/pkg/handler"
)

func SetupRoutes(root fiber.Router) {
	root.Get("hello-world", handler.HelloWorld())
}
