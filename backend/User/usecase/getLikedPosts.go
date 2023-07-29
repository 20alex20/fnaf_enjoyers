package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"net/http"
)

func (uc *useCase) GetLikedPosts(nickname string, repo repository.Repository) ([]model.PostUser, error) {
	postIDs, err := repo.GetLikedPostsIDs(nickname)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://localhost:3001/post/ids")

	jsn, err := json.Marshal(postIDs)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsn)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var posts []model.PostUser
	if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}
