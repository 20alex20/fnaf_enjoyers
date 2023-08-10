package repository

const Post = `
	select * from post where id = $1;
`

const RejectedPost = `
	select id, "date", "text" from post 
	where user_id = $1 
	and checked = true 
	and accepted = false
	order by date desc;
`

const RejectedMessage = `
	select text from message where post_id = $1;
`

const CheckingPost = `
	select id, user_id, "date", "text"
	from post
	where user_id = $1
	and checked = false;
`

const PostCategories = `
	select title from category where id in (
		select category_id from post_category
			where post_id = $1
	);
`

const PostFilters = `
	select title from filter where id in (
		select filter_id from post_filter
			where post_id = $1
	);
`

const Liked = `
	select * from post 
	where id in (
		select post_id from user_post_liked
		where user_id = $1
	);
`

const CreatePostReturnID = `
	insert into post
			(user_id, text)
		values (
		        $1,
		 		$2
		 )
		returning id;
`

const CreateCategoryRelation = `
	insert into post_category
				(post_id, category_id)
			values (	
			 	$1, 
				(select id from category where title=$2)
			);
`

const CreateFilterRelation = `
	insert into post_filter
				(post_id, filter_id)
			values (	
			 	$1, 
				(select id from filter where title=$2)
			);
`

const UpdateInfoTotalPosts = `
	update info
		set total_posts=total_posts + 1;
`

const GetIDsByTitles = `
	select id from $1 where title in $2
`

const GetUserPosts = `
	select * from post where user_id = $1
		order by date desc;
`
