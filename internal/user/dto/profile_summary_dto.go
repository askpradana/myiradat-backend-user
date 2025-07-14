package dto

import "time"

type GetProfileSummaryRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type CommentDTO struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
}

type ConsultDTO struct {
	ConsultDate          *time.Time   `json:"consultDate"`
	LatestAnalysisResult string       `json:"latestAnalysisResult"`
	LatestComments       []CommentDTO `json:"latestComments"`
}

type TestResultDTO[T any] struct {
	TestTakenDate *time.Time `json:"testTakenDate"`
	Result        T          `json:"result"`
}

type TestsDTO struct {
	Ipro  TestResultDTO[IproTestResult]  `json:"ipro"`
	Iprob TestResultDTO[IproTestResult]  `json:"iprob"`
	Ipros TestResultDTO[IprosTestResult] `json:"ipros"`
}

type ProfileInfoDTO struct {
	Email string `json:"email"`
	NoHP  string `json:"no_hp"`
	Name  string `json:"name"`
}

type ProfileSummaryResponse struct {
	Profile  ProfileInfoDTO `json:"profile"`
	Consults ConsultDTO     `json:"consults"`
	Tests    TestsDTO       `json:"tests"`
}
