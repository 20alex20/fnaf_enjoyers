package repository

func (r *repository) ChangeNickname(old, new string) error {
	query := `
		update "user"
		set nickname=$1
		where id=(
		    select id
		    from "user"
		    where nickname=$2
		);
	`

	_, err := r.DB.Exec(query, new, old)
	return err
}
