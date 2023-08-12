package route

import (
	"github.com/fnaf-enjoyers/user-service/handler"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app fiber.Router, uc usecase.UseCase, repo repository.Repository) {
	app.Get("nickname", handler.GetNickname(repo))
	app.Get("profile-pic", handler.GetProfilePic(repo))

	user := app.Group("user")
	user.Get("nickname", handler.GetUserNickname())
	user.Post("nickname", handler.ChangeNickname(uc, repo))
	user.Post("register", handler.RegisterUser(uc, repo))
	user.Post("auth", handler.AuthUser(uc, repo))
	user.Post("logout", handler.LogOut())
	user.Get("role", handler.GetUserRole(uc, repo))
	user.Get("exist", handler.CheckNickname(uc, repo))
	user.Get("profile-pic", handler.GetUserProfilePic(repo))
	user.Post("profile-pic", handler.ChangeProfilePic(repo))

	app.Get("posts", handler.GetUserPosts(repo))
	posts := app.Group("posts")
	posts.Get("liked", handler.GetLikedPosts(repo))
	posts.Get("rejected", handler.GetRejectedPosts(repo))
	posts.Get("moder", handler.GetModerPosts(repo))

	app.Get("post", handler.GetPost(uc, repo))
	post := app.Group("post")
	post.Post("create", handler.CreatePost(repo))
	post.Post("verify", handler.VerifyPost(uc, repo))
	post.Post("reject", handler.RejectPost(repo))
	post.Get("check-like", handler.CheckLike(repo))
	post.Post("set-like", handler.SetLike(uc, repo))
	post.Post("unset-like", handler.UnsetLike(uc, repo))
	post.Post("comment", handler.LeftComment(uc, repo))
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
