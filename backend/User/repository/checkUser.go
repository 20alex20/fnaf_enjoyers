package repository

func (r *repository) CheckUser(nickname string) (bool, error) {
	var exist bool
	query := `
		select exists(select 1 from "user" where nickname = $1) as "exist";
	`

	err := r.DB.Get(&exist, query, nickname)
	if err != nil {
		return false, err
	}

	return exist, nil
}
