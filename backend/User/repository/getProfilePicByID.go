package repository

func (r *repository) GetProfilePicByID(userID string) (string, error) {
	link := ""

	query := `
		select link 
		from profile_picture 
		where id = (
		    select profile_picture_id from "user" where id = $1
		);
`

	err := r.DB.Get(&link, query, userID)
	if err != nil {
		return "", err
	}

	return link, nil
}
