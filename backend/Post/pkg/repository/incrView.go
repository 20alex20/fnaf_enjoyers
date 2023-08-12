package repository

func (r *repository) IncrView(postID string) error {
	_, err := r.DB.Exec(IncrPostViews, postID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(UpdateInfoTotalViews)
	return err
}
