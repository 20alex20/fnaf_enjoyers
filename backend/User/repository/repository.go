package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
	CheckUser(nickname string) (bool, error)
	CreateUser(nickname, password string) error
	GetPasswordHash(nickname string) (string, error)
	ChangeNickname(old, new string) error
}

type repository struct {
	DB sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		DB: *db,
	}
}
