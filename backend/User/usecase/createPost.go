package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"net/http"
)

func (uc *useCase) CreatePost(nickname, text string, categories, filters []string, repo repository.Repository) error {
	url := fmt.Sprintf("http://localhost:3001/post/create?nickname=%s", nickname)
	body := model.CreatePost{
		Categories: categories,
		Filters:    filters,
		Text:       text,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	err = repo.IncrPosts(nickname)
	return err
}
