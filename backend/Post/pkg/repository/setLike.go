package repository

func (r *repository) SetLike(postID string) error {
	_, err := r.DB.Exec(IncrPostLikes, postID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(UpdateInfoTotalLikes)
	return err
}
