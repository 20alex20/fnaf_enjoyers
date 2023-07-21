package route

import (
	"github.com/fnaf-enjoyers/post-service/pkg/handler"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecases"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app fiber.Router, uc usecases.UseCase, repo repository.Repository) {
	main := app.Group("main")
	main.Get("posts", handler.MainPosts(uc, repo))

	app.Get("user-post", handler.UserPost(uc, repo))
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
