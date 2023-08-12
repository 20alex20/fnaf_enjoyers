package repository

func (r *repository) GetNicknameByID(userID string) (string, error) {
	nickname := ""

	query := `select nickname from "user" where id = $1`

	err := r.DB.Get(&nickname, query, userID)
	if err != nil {
		return "", err
	}

	return nickname, nil
}
