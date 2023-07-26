package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type NicknameResponse struct {
	Nickname string `json:"nickname"`
}
