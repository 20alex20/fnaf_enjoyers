package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetRejectedPosts(userID string, repo repository.Repository) ([]model.RejectedPostResponse, error) {
	rejectedPostsDTO, err := repo.GetRejectedPosts(userID)
	if err != nil {
		return nil, err
	}

	var rejectedPosts []model.RejectedPostResponse

	for _, rejectedPostDTO := range rejectedPostsDTO {
		categories, err := repo.GetPostCategories(rejectedPostDTO.ID)
		if err != nil {
			return nil, err
		}

		filters, err := repo.GetPostFilters(rejectedPostDTO.ID)
		if err != nil {
			return nil, err
		}

		moderText, err := repo.GetRejectedMessage(rejectedPostDTO.ID)
		if err != nil {
			return nil, err
		}

		rejectedPost := model.RejectedPostResponse{
			ID:         rejectedPostDTO.ID,
			Text:       rejectedPostDTO.Text,
			DateTime:   rejectedPostDTO.Date.Format("2006.01.02 15:04"),
			Categories: categories,
			Filters:    filters,
			ModerText:  moderText,
		}

		rejectedPosts = append(rejectedPosts, rejectedPost)
	}

	return rejectedPosts, nil
}
