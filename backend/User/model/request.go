package model

type RegisterUserRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type CreatePost struct {
	Categories []string `json:"categories"`
	Filters    []string `json:"filters"`
	Text       string   `json:"text"`
}
