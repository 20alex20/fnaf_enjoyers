package repository

func (r *repository) VerifyPost(postID string) error {
	_, err := r.DB.Exec(VerifyPost, postID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(UpdateInfoTotalPosts)
	return err
}
