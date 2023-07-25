package repository

func (r *repository) GetPasswordHash(nickname string) (string, error) {
	var password string
	query := `
		select password from "user" where nickname = $1
	`
	err := r.DB.Get(&password, query, nickname)
	if err != nil {
		return "", err
	}

	return password, nil
}
