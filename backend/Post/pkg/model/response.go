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
	Categories []string `json:"categories" example:"['IU', 'IU-5']"`
	Filters    []string `json:"filters" example:"['funny']"`
	ModerText  string   `json:"moder_text" example:"There is a name of student in the post"`
}

type CheckingPostResponse struct {
	ID         string   `json:"id" example:"some-uuid"`
	Text       string   `json:"text" example:"Lorem ipsum"`
	DateTime   string   `json:"date_time" example:"2023.07.20 12:00"`
	Categories []string `json:"categories" example:"['IU', 'IU-5']"`
	Filters    []string `json:"filters" example:"['funny']"`
	UserID     string   `json:"nickname" example:"John Doe"`
}
