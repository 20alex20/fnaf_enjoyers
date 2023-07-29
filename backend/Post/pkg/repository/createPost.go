package repository

func (r *repository) CreatePost(categories, filters []string, text, nickname string) error {
	query := `
		insert into post
			(user_id, text)
		values (
		        (select id from "user" where nickname=$1),
		 		$2
		 )
		returning id;
	`

	var postID string
	err := r.DB.Get(&postID, query, nickname, text)
	if err != nil {
		return err
	}

	for _, category := range categories {
		query = `
			insert into post_category
				(post_id, category_id)
			values (	
			 	$1, 
				(select id from category where title=$2)
			);
		`

		_, err = r.DB.Exec(query, postID, category)
		if err != nil {
			return err
		}
	}

	for _, filter := range filters {
		query = `
			insert into post_filter
				(post_id, filter_id)
			values (	
			 	$1, 
				(select id from filter where title=$2)
			);
		`

		_, err = r.DB.Exec(query, postID, filter)
		if err != nil {
			return err
		}
	}

	query = `
		update info
		set total_posts=total_posts + 1;
	`
	_, err = r.DB.Exec(query)

	return err
}
