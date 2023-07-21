package model

type MainPostRequest struct {
	Category string `json:"category" example:"all"`
	Filter   string `json:"filter" example:"without"`
	Sort     string `json:"sort" enums:"date_time,likes,views"`
	Number   int    `json:"number" example:"2"`
	Page     int    `json:"page" example:"1"`
}
