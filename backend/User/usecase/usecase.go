package usecase

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
)

type UseCase interface {
	RegisterUser(req model.RegisterUserRequest, repo repository.Repository) error
	AuthUser(req model.RegisterUserRequest, repo repository.Repository) error
	ChangeNickname(old, new string, repo repository.Repository) error
	CreatePost(nickname, text string, categories, filters []string, repo repository.Repository) error
	GetUserPosts(nickname string) ([]model.PostUser, error)
	GetLikedPosts(nickname string, repo repository.Repository) ([]model.PostUser, error)
	GetUserRole(nickname string, repo repository.Repository) (string, error)
}

type useCase struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) UseCase {
	return &useCase{
		cfg: cfg,
	}
}
