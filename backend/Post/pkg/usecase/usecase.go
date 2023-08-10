package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/config"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

type UseCase interface {
	MainPosts(category, filter, sort string, number, page int, repo repository.Repository) ([]model.PostResponse, error)
	ProcessPostDTO(postsDTO []model.PostDTO) (posts []model.PostResponse)
	GetUserPosts(nickname string, repo repository.Repository) ([]model.PostResponse, error)
	GetRejectedPosts(userID string, repo repository.Repository) ([]model.RejectedPostResponse, error)
	GetLikedPosts(userID string, repo repository.Repository) ([]model.PostResponse, error)
	GetCheckingPosts(userID string, repo repository.Repository) ([]model.CheckingPostResponse, error)
	GetPost(postID string, repo repository.Repository) (model.PostResponse, error)
}

type useCase struct {
	cfg config.Config
}

func NewService(cfg *config.Config) UseCase {
	return &useCase{
		cfg: *cfg,
	}
}
