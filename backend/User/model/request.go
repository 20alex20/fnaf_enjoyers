package model

type UserRequest struct {
	Nickname string `json:"nickname" example:"John Doe"`
	Password string `json:"password" example:"john-qwerty"`
}

type CreatePostRequest struct {
	Categories []string `json:"categories[]"`
	Filters    []string `json:"filters[]"`
	Text       string   `json:"text" example:"Lorem ipsum"`
}
