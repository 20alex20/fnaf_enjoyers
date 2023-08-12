package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetUserPosts(userID string, repo repository.Repository) ([]model.PostResponse, error) {
	postsDTO, err := repo.GetPostByUserID(userID)
	if err != nil {
		return nil, err
	}

	posts := uc.ProcessPostDTO(postsDTO)

	return posts, nil
}
