package model

type PostMain struct {
	Text     string `json:"text" example:"Lorem ipsum"`
	DateTime string `json:"date_time" example:"2023.07.20 12:00"`
	Views    int    `json:"views" example:"100"`
	Likes    int    `json:"likes" example:"50"`
}

type PostUser struct {
	ID       int    `json:"id" example:"1"`
	Text     string `json:"text" example:"Lorem ipsum"`
	DateTime string `json:"date_time" example:"2023.07.20 12:00"`
	Views    int    `json:"views" example:"100"`
	Likes    int    `json:"likes" example:"50"`
}
