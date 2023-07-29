package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/user-service/model"
	"io"
	"net/http"
)

func (uc *useCase) GetUserPosts(nickname string) ([]model.PostUser, error) {
	url := fmt.Sprintf("http://localhost:3001/post/get?nickname=%s", nickname)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(res.Body)

	var posts []model.PostUser
	if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}
