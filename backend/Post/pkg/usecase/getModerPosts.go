package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"net/http"
)

func (uc *useCase) GetModerPosts(repo repository.Repository) ([]model.ModerPostResponse, error) {
	moderPostsDTO, err := repo.GetPostsForModerator()
	if err != nil {
		return nil, err
	}

	var moderPosts []model.ModerPostResponse

	for _, moderPostDTO := range moderPostsDTO {
		categories, err := repo.GetPostCategories(moderPostDTO.ID)
		if err != nil {
			return nil, err
		}

		filters, err := repo.GetPostFilters(moderPostDTO.ID)
		if err != nil {
			return nil, err
		}

		url := fmt.Sprintf("http://localhost:3002/nickname?id=%s", moderPostDTO.UserID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		nickname := ""
		if err = json.NewDecoder(res.Body).Decode(&nickname); err != nil {
			return nil, err
		}

		moderPost := model.ModerPostResponse{
			ID:         moderPostDTO.ID,
			Text:       moderPostDTO.Text,
			DateTime:   moderPostDTO.Date.Format("2006.01.02 15:04"),
			Categories: categories,
			Filters:    filters,
			Nickname:   nickname,
		}

		moderPosts = append(moderPosts, moderPost)
		_ = res.Body.Close()
	}

	return moderPosts, nil
}
