package dto

import "time"

type CreateIprobTestRequest struct {
	TestID        *int           `json:"testId"`
	Result        IproTestResult `json:"result" binding:"required"`
	TestTakenDate *time.Time     `json:"testTakenDate"`
	Email         string         `json:"email" binding:"required,email"`
}
