package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
	CheckUser(nickname string) (bool, error)
	CreateUser(nickname, password string) error
	GetPasswordHash(nickname string) (string, error)
	ChangeNickname(old, new string) error
	IncrPosts(userID string) error
	GetUserRole(nickname string) (*bool, error)
	GetUserID(nickname string) (string, error)
	GetNicknameByID(userID string) (string, error)
	CheckIfModerator(nickname string) (bool, error)
	CheckLike(userID, postID string) (bool, error)
	SetLike(userID, postID string) error
	IncrLikes(userID string) error
	IncrViews(userID string) error
	UnsetLike(userID, postID string) error
	DecrLikes(userID string) error
	GetProfilePicByID(userID string) (string, error)
	ChangeProfilePic(userID, link string) error
}

type repository struct {
	DB sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		DB: *db,
	}
}
