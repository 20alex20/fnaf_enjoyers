package model

type PostResponse struct {
	ID       string `json:"id" example:"some-uuid"`
	Text     string `json:"text" example:"Lorem ipsum"`
	DateTime string `json:"date_time" example:"2023.07.20 12:00"`
	Views    int    `json:"views" example:"100"`
	Likes    int    `json:"likes" example:"50"`
}

type RejectedPostResponse struct {
	ID         string   `json:"id" example:"some-uuid"`
	Text       string   `json:"text" example:"Lorem ipsum"`
	DateTime   string   `json:"date_time" example:"2023.07.20 12:00"`
	Categories []string `json:"categories"`
	Filters    []string `json:"filters"`
	ModerText  string   `json:"moder_text" example:"There is a student's name in post"`
}

type ModerPostResponse struct {
	ID         string   `json:"id" example:"some-uuid"`
	Text       string   `json:"text" example:"Lorem ipsum"`
	DateTime   string   `json:"date_time" example:"2023.07.20 12:00"`
	Categories []string `json:"categories"`
	Filters    []string `json:"filters"`
	Nickname   string   `json:"nickname" example:"John Doe"`
}

type CheckUsernameResponse struct {
	ThereIs bool `json:"there_is"`
}

type CheckLikeResponse struct {
	State string `json:"state"`
}

type ImgResponse struct {
	Img string `json:"img"`
}
