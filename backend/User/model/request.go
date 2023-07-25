package model

type RegisterUserRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
