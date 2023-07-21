package model

import "time"

type PostDTO struct {
	ID       string    `db:"id"`
	UserID   string    `db:"user_id"`
	Date     time.Time `db:"date"`
	Text     string    `db:"text"`
	Views    int       `db:"views"`
	Likes    int       `db:"likes"`
	Accepted bool      `db:"accepted"`
}

type CategoryDTO struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

type FilterDTO struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

type PostFilterDTO struct {
	PostID   string `db:"post_id"`
	FilterID string `db:"filter_id"`
}

type PostCategoryDTO struct {
	PostID     string `db:"post_id"`
	CategoryID string `db:"category_id"`
}
