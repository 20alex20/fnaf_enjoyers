package usecases

import (
	"github.com/fnaf-enjoyers/post-service/pkg/config"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

type UseCase interface {
	UserPost(userID string, repo repository.Repository) ([]model.Post, error)
	MainPosts(requestBody model.MainPostRequest, repo repository.Repository) ([]model.PostMain, error)
}

type service struct {
	cfg config.Config
}

func NewService(cfg *config.Config) UseCase {
	return &service{
		cfg: *cfg,
	}
}
