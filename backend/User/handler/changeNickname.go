package handler

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
)

// ChangeNickname godoc
// @Summary Change nickname
// @Description Change nickname for current user
// @Tags user
// @Accept json
// @Param req body model.NicknameJSON true "request body"
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 400 {string} string
// @Success 200 {string} string
// @Router /user/nickname [post]
func ChangeNickname(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		newNickname := ctx.FormValue("nickname", "")
		if newNickname == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("new nickname is empty")
		}

		err = uc.ChangeNickname(nickname.(string), newNickname, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		session.Set("name", newNickname)
		err = session.Save()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}
