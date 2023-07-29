package model

type RegisterUserRequest struct {
	Nickname string `json:"nickname" example:"John Doe"`
	Password string `json:"password" example:"john-qwerty"`
}

type CreatePost struct {
	Categories []string `json:"categories"`
	Filters    []string `json:"filters"`
	Text       string   `json:"text" example:"Lorem ipsum"`
}
