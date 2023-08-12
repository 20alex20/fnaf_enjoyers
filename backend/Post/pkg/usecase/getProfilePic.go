package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *useCase) GetProfilePic(userID string) (string, error) {
	url := fmt.Sprintf("http://localhost:3002/profile-pic?id=%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	profilePic := ""
	if err = json.NewDecoder(res.Body).Decode(&profilePic); err != nil {
		return "", err
	}

	return profilePic, nil
}
