package model

import "time"

type PostDTO struct {
	ID       string    `db:"id"`
	UserID   string    `db:"user_id"`
	Date     time.Time `db:"date"`
	Text     string    `db:"text"`
	Views    int       `db:"views"`
	Likes    int       `db:"likes"`
	Checked  bool      `db:"checked"`
	Accepted bool      `db:"accepted"`
}

type RejectedPostDTO struct {
	ID   string    `db:"id"`
	Date time.Time `db:"date"`
	Text string    `db:"text"`
}

type CheckingPostDTO struct {
	ID     string    `db:"id"`
	UserID string    `db:"user_id"`
	Date   time.Time `db:"date"`
	Text   string    `db:"text"`
}
