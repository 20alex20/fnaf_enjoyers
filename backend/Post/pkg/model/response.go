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

type ModerPostResponse struct {
	ID         string   `json:"id" example:"some-uuid"`
	Text       string   `json:"text" example:"Lorem ipsum"`
	DateTime   string   `json:"date_time" example:"2023.07.20 12:00"`
	Categories []string `json:"categories"`
	Filters    []string `json:"filters"`
	Nickname   string   `json:"nickname" example:"John Doe"`
}

type CommentResponse struct {
	ID      string            `json:"id" example:"some-uuid"`
	Text    string            `json:"text" example:"Lorem ipsum"`
	Date    string            `json:"date" example:"2023.07.20 12:00"`
	Author  string            `json:"author" example:"John Doe"`
	Img     string            `json:"img" example:"/path/to/image"`
	Replies []CommentResponse `json:"replies,omitempty"`
}

type MaxPageResponse struct {
	MaxPage int `json:"max_page"`
}
