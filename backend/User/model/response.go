package model

type PostResponse struct {
	ID       string `json:"id" example:"some-uuid"`
	Text     string `json:"text" example:"Lorem ipsum"`
	DateTime string `json:"date_time" example:"2023.07.20 12:00"`
	Views    int    `json:"views" example:"100"`
	Likes    int    `json:"likes" example:"50"`
}

type CheckUsernameResponse struct {
	ThereIs bool `json:"there_is"`
}
