package usecase

import (
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (uc *useCase) RegisterUser(req model.UserRequest, repo repository.Repository) error {
	exist, err := repo.CheckUser(req.Nickname)
	if err != nil {
		return errors.Errorf("Не удалось проверить уникальность имени пользователя")
	}

	if exist {
		return errors.Errorf(
			"Пользователь с ником %s уже существует!",
			req.Nickname)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return err
	}

	err = repo.CreateUser(req.Nickname, string(passwordHash))

	return nil
}
