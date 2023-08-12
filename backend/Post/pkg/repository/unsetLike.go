package repository

func (r *repository) UnsetLike(postID string) error {
	_, err := r.DB.Exec(DecrPostLikes, postID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(DecrInfoTotalLikes)
	return err
}
