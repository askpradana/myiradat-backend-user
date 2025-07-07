package dto

type RequestGetDetail struct {
}

type ResponseGetDetail struct {
	UserID   string `json:"userId" example:"12345"`
	Username string `json:"username" example:"vicky"`
}
