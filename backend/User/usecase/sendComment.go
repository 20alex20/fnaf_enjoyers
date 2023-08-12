package usecase

import (
	"net/http"
	"net/url"
	"strings"
)

func (uc *useCase) SendComment(text, userID, referenceID string) error {
	commentUrl := "http://localhost:3001/post/comment"
	formData := make(url.Values)
	formData.Add("text", text)
	formData.Add("reference_id", referenceID)
	formData.Add("user_id", userID)
	req, err := http.NewRequest("POST", commentUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err = http.DefaultClient.Do(req)
	return err
}
