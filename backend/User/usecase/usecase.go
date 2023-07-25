package usecase

import (
	"github.com/fnaf-enjoyers/user/config"
	"github.com/fnaf-enjoyers/user/model"
	"github.com/fnaf-enjoyers/user/repository"
)

type UseCase interface {
	RegisterUser(req model.RegisterUserRequest, repo repository.Repository) error
	AuthUser(req model.RegisterUserRequest, repo repository.Repository) error
}

type useCase struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) UseCase {
	return &useCase{
		cfg: cfg,
	}
}
