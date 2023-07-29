package usecase

import "github.com/fnaf-enjoyers/post-service/pkg/repository"

func (uc *useCase) CreatePost(categories, filters []string, text, nickname string, repo repository.Repository) error {
	err := repo.CreatePost(categories, filters, text, nickname)
	return err
}
