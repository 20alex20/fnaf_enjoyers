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

const GetUserPosts = `
	select * from post 
	         where user_id = $1
	         and checked = true
	         and accepted = true
		order by date desc;
`

const GetModerPosts = `
	select * from post where checked = false and accepted = false
		order by date desc;
`

const VerifyPost = `
	update post
	set checked = true, accepted = true
	where id = $1;
`

const RejectPost = `
	update post
	set checked = true, accepted = false
	where id = $1;
`

const InsertRejectedMessage = `
	insert into message
	(post_id, text)
	values ($1, $2);
`

const IncrPostLikes = `
	update post
	set likes = likes + 1
	where id = $1;
`

const DecrPostLikes = `
	update post
	set likes = likes - 1
	where id = $1;
`

const IncrPostViews = `
	update post
	set views = views + 1
	where id = $1;
`

const UpdateInfoTotalLikes = `
	update info
		set total_likes = total_likes + 1;
`

const DecrInfoTotalLikes = `
	update info
		set total_likes = total_likes - 1;
`

const UpdateInfoTotalViews = `
	update info
		set total_views = total_views + 1;
`

const GetPostAuthor = `
	select user_id from post where id = $1;
`

const GetComments = `
	select * from comment where reference_id = $1 
	order by date desc;
`

const LeftComment = `
	insert into comment
	(reference_id, user_id, text)
	values ($1, $2, $3);
`

const CountMainPosts = `
	select count(*) from post where checked = true and accepted = true;
`

const CountMainPostsCategorized = `
	select count(*) from post 
		         where id in 
		               (select post_id from post_category where category_id in (
		               		select id from category where title = $1 
					   ))
				 and checked = true
				 and accepted = true;
`

const CountMainPostsFiltered = `
	select count(*) from post 
		         where id in 
		               (select post_id from post_filter where filter_id in (
		                   select id from filter where title = $1
		               ))
				 and checked = true
				 and accepted = true;
`

const CountMainPostsCF = `
	select * from post 
	where id in (
	    select pf.post_id 
		from post_filter pf 
		    inner join post_category pc 
		        on pc.post_id = pf.post_id 
		where 
		    pf.filter_id in
			(select id from filter where title = $1)
		and 
		    pc.category_id in
			(select id from category where title = $2)
	) 
	  and checked=true 
	  and accepted=true;
`
