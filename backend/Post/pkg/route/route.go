package route

import (
	"github.com/fnaf-enjoyers/post-service/pkg/handler"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app fiber.Router, uc usecase.UseCase, repo repository.Repository) {
	main := app.Group("main")
	main.Get("posts", handler.MainPosts(uc, repo))

	post := app.Group("post")
	post.Post("create", handler.CreatePost(uc, repo))
	post.Get("get", handler.GetUserPosts(uc, repo))
	post.Post("ids", handler.GetPostsByIDs(uc, repo))
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
