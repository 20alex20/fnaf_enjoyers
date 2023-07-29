package usecase

import (
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (uc *useCase) AuthUser(req model.RegisterUserRequest, repo repository.Repository) error {
	exist, err := repo.CheckUser(req.Nickname)
	if err != nil {
		return err
	}

	if !exist {
		return fiber.ErrNotFound
	}

	passwordHash, err := repo.GetPasswordHash(req.Nickname)
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
}
