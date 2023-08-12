package repository

func (r *repository) IncrLikes(userID string) error {
	query := `
		update "user"
		set user_likes = user_likes + 1
		where id=$1;
`
	_, err := r.DB.Exec(query, userID)
	return err
}
