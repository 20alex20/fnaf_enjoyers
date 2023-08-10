package route

import (
	"github.com/fnaf-enjoyers/user-service/handler"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app fiber.Router, uc usecase.UseCase, repo repository.Repository) {
	user := app.Group("user")
	user.Get("nickname", handler.GetNickname())
	user.Post("nickname", handler.ChangeNickname(uc, repo))
	user.Post("register", handler.RegisterUser(uc, repo))
	user.Post("auth", handler.AuthUser(uc, repo))
	user.Post("logout", handler.LogOut())
	user.Get("role", handler.GetUserRole(uc, repo))
	user.Post("exist", handler.CheckNickname(uc, repo))

	post := app.Group("post")
	post.Post("create", handler.CreatePost(repo))
	post.Get("get", handler.GetUserPosts(repo))
	post.Get("liked", handler.GetLikedPosts(repo))
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
