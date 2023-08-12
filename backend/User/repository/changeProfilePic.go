package repository

func (r *repository) ChangeProfilePic(userID, link string) error {
	var picID string
	query := `
		insert into profile_picture
		(link) values ($1)
		returning id;
`

	err := r.DB.Get(&picID, query, link)
	if err != nil {
		return err
	}

	query = `
		update "user"
		set profile_picture_id = $1
		where id = $2;
`
	_, err = r.DB.Exec(query, picID, userID)
	if err != nil {
		return err
	}

	return nil
}
