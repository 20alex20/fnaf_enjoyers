package handler

import (
	"fmt"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
)

// GetProfilePic godoc
// @Summary Get nickname
// @Description Get nickname by user id.
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Success 200 {string} string
// @Router /nickname [get]
func GetProfilePic(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("id", "")

		if userID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("empty id")
		}

		link, err := repo.GetProfilePicByID(userID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		profilePic := fmt.Sprintf("profile_pics/%s", link)

		return ctx.Status(fiber.StatusOK).JSON(profilePic)
	}
}
