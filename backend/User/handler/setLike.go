package handler

import (
	"fmt"
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SetLike(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
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

		postID := ctx.FormValue("post_id", "")
		if postID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("post_id is empty")
		}

		err = repo.SetLike(userID, postID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		url := fmt.Sprintf("http://localhost:3001/post/set-like?post_id=%s", postID)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		authorID, err := uc.GetPostAuthor(postID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		err = repo.IncrLikes(authorID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		ctx.Status(fiber.StatusOK)
		return nil
	}
}
