package model

type Post struct {
	Username string
	Date     string
	Text     string
	Views    int
	Likes    int
	Comments int
}

type PostMain struct {
	Text     string `json:"text" example:"Lorem ipsum"`
	DateTime string `json:"date_time" example:"2023.07.20 12:00"`
	Views    int    `json:"views" example:"100"`
	Likes    int    `json:"likes" example:"50"`
}
