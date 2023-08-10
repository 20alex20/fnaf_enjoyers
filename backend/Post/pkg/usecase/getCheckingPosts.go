package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetCheckingPosts(userID string, repo repository.Repository) ([]model.CheckingPostResponse, error) {
	checkingPostsDTO, err := repo.GetCheckingPosts(userID)
	if err != nil {
		return nil, err
	}

	var checkingPosts []model.CheckingPostResponse

	for _, checkingPostDTO := range checkingPostsDTO {
		categories, err := repo.GetPostCategories(checkingPostDTO.ID)
		if err != nil {
			return nil, err
		}

		filters, err := repo.GetPostFilters(checkingPostDTO.ID)
		if err != nil {
			return nil, err
		}

		checkingPost := model.CheckingPostResponse{
			ID:         checkingPostDTO.ID,
			Text:       checkingPostDTO.Text,
			DateTime:   checkingPostDTO.Date.Format("2006.01.02 15:04"),
			Categories: categories,
			Filters:    filters,
			UserID:     checkingPostDTO.UserID,
		}

		checkingPosts = append(checkingPosts, checkingPost)
	}

	return checkingPosts, nil
}
