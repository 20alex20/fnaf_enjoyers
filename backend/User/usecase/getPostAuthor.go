package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *useCase) GetPostAuthor(postID string) (string, error) {
	url := fmt.Sprintf("http://localhost:3001/post/author?post_id=%s", postID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var authorID string
	if err = json.NewDecoder(res.Body).Decode(&authorID); err != nil {
		return "", nil
	}

	return authorID, nil
}
