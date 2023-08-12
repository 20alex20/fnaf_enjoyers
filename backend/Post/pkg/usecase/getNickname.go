package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *useCase) GetNickname(userID string) (string, error) {
	url := fmt.Sprintf("http://localhost:3002/nickname?id=%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	nickname := ""
	if err = json.NewDecoder(res.Body).Decode(&nickname); err != nil {
		return "", err
	}

	return nickname, nil
}
