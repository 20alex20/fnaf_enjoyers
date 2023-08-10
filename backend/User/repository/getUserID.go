package repository

func (r *repository) GetUserID(nickname string) (string, error) {
	var userID string

	query := `
		select id from "user" where nickname = $1;
`
	err := r.DB.Get(&userID, query, nickname)
	if err != nil {
		return "", err
	}

	return userID, nil
}
