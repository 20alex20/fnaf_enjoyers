package repository

func (r *repository) CheckIfModerator(nickname string) (bool, error) {
	var isModerator bool
	query := `select is_moderator from "user" where nickname = $1`

	err := r.DB.Get(&isModerator, query, nickname)
	if err != nil {
		return false, err
	}

	return isModerator, nil
}
