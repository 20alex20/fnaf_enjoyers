package usecase

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
)

type UseCase interface {
	RegisterUser(req model.UserRequest, repo repository.Repository) error
	AuthUser(req model.UserRequest, repo repository.Repository) error
	ChangeNickname(old, new string, repo repository.Repository) error
	GetUserRole(nickname string, repo repository.Repository) (string, error)
	CheckNickname(nickname string, repo repository.Repository) (bool, error)
	GetPostAuthor(postID string) (string, error)
	SendComment(text, userID, referenceID string) error
}

type useCase struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) UseCase {
	return &useCase{
		cfg: cfg,
	}
}
