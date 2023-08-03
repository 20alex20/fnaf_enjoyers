package model

type NicknameJSON struct {
	Nickname string `json:"nickname" example:"John Doe"`
}

type UserRole struct {
	Status string `json:"status"`
}
