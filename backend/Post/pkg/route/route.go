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
	main.Get("max-page", handler.MaxPage(uc, repo))

	app.Get("posts", handler.GetUserPosts(uc, repo))
	posts := app.Group("posts")
	posts.Get("liked", handler.GetLikedPosts(uc, repo))
	posts.Get("rejected", handler.GetRejectedPosts(uc, repo))
	posts.Get("moder", handler.GetModerPosts(uc, repo))

	app.Get("post", handler.GetPost(uc, repo))
	post := app.Group("post")
	post.Post("create", handler.CreatePost(repo))
	post.Post("verify", handler.VerifyPost(repo))
	post.Post("reject", handler.RejectPost(repo))
	post.Post("set-like", handler.SetLike(repo))
	post.Post("unset-like", handler.UnsetLike(repo))
	post.Post("incr-view", handler.IncrView(repo))
	post.Get("author", handler.GetPostAuthor(repo))
	post.Get("comments", handler.GetComments(uc, repo))
	post.Post("comment", handler.LeftComment(repo))
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
