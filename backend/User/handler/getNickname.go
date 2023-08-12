package handler

import (
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
)

// GetNickname godoc
// @Summary Get nickname
// @Description Get nickname by user id.
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Success 200 {string} string
// @Router /nickname [get]
func GetNickname(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("id", "")

		if userID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("empty id")
		}

		nickname, err := repo.GetNicknameByID(userID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(nickname)
	}
}
