package route

import (
	"github.com/fnaf-enjoyers/user/handler"
	"github.com/fnaf-enjoyers/user/repository"
	"github.com/fnaf-enjoyers/user/usecase"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app fiber.Router, uc usecase.UseCase, repo repository.Repository) {
	user := app.Group("user")
	user.Post("register", handler.RegisterUser(uc, repo))
	user.Post("auth", handler.AuthUser(uc, repo))
	user.Post("logout", handler.LogOut())
	user.Get("whoami", handler.WhoAmI())
	user.Get("nickname", handler.GetNickname())
}

func SetupSwagger(app fiber.Router) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
