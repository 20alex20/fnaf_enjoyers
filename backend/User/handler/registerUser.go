package handler

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
)

// RegisterUser godoc
// @Summary Register new user
// @Description Creates new user and processes authentication.
// @Tags user
// @Accept json
// @Param req body model.RegisterUserRequest true "request body"
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 400 {string} string
// @Success 200 {string} string
// @Router /user/register [post]
func RegisterUser(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req model.RegisterUserRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		err = uc.RegisterUser(req, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		session.Set("name", req.Nickname)
		err = session.Save()
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
		}

		ctx.Status(fiber.StatusOK)

		return nil
	}
}
