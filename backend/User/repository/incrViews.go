package repository

func (r *repository) IncrViews(userID string) error {
	query := `
		update "user"
		set user_views = user_views + 1
		where id=$1;
`
	_, err := r.DB.Exec(query, userID)
	return err
}
