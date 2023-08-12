package handler

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
)

func CheckLike(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		userID, err := repo.GetUserID(nickname.(string))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		postID := ctx.Query("post_id", "")
		if postID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("post_id is empty")
		}

		liked, err := repo.CheckLike(userID, postID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		res := model.CheckLikeResponse{
			State: "not_clicked",
		}

		if liked {
			res.State = "clicked"
		}

		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}
