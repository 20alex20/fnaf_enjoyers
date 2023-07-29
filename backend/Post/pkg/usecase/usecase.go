package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/config"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

type UseCase interface {
	MainPosts(category, filter, sort string, number, page int, repo repository.Repository) ([]model.PostMain, error)
	CreatePost(categories, filters []string, text, nickname string, repo repository.Repository) error
	GetUserPosts(nickname string, repo repository.Repository) ([]model.PostUser, error)
}

type useCase struct {
	cfg config.Config
}

func NewService(cfg *config.Config) UseCase {
	return &useCase{
		cfg: *cfg,
	}
}
