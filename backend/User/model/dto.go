package model

type User struct {
	ID               string `db:"id"`
	ProfilePictureID string `db:"profile_picture_id"`
	Nickname         string `db:"nickname"`
	Password         string `db:"password"`
	IsModerator      bool   `db:"is_moderator"`
	UserPosts        int    `db:"user_posts"`
	UserLikes        int    `db:"user_likes"`
	UserViews        int    `db:"user_views"`
}

type Exist struct {
	Exist bool `db:"exist"`
}
