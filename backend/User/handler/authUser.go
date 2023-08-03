package handler

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
)

// AuthUser godoc
// @Summary Authenticate user
// @Description Checks user credentials and logs in if successful.
// @Tags user
// @Accept json
// @Param req body model.RegisterUserRequest true "request body"
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 400 {string} string
// @Success 200 {string} string
// @Router /user/auth [post]
func AuthUser(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req model.RegisterUserRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		err = uc.AuthUser(req, repo)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
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
