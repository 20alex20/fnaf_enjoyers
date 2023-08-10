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

	app.Get("posts", handler.GetUserPosts(uc, repo))
	app.Get("post", handler.GetPost(uc, repo))
	post := app.Group("post")
	post.Post("create", handler.CreatePost(repo))
	post.Get("rejected", handler.GetRejectedPosts(uc, repo))
	post.Get("liked", handler.GetLikedPosts(uc, repo))
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
