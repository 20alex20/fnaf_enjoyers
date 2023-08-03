package repository

func (r *repository) GetUserRole(nickname string) (*bool, error) {
	var isModerator bool
	query := `
		select is_moderator from "user" where nickname = $1;
`

	err := r.DB.Get(&isModerator, query, nickname)
	if err != nil {
		return nil, err
	}

	return &isModerator, nil
}
