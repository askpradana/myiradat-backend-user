package dto

import "time"

type CommentRequest struct {
	Comment string `json:"comment"`
}

type CreateConsultRequest struct {
	Email          string           `json:"email" binding:"required,email"`
	Owner          string           `json:"owner"`
	ConsultDate    *time.Time       `json:"consultDate"`
	AnalysisResult string           `json:"analysisResult"`
	Comments       []CommentRequest `json:"comments"`
}
