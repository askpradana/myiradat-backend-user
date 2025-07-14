package dto

import "time"

type IprosTestResult struct {
	HubunganInterpersonal string `json:"hubunganInterpersonal"`
	KecepatanPerseptual   string `json:"kecepatanPerseptual"`
	KecerdasanUmum        string `json:"kecerdasanUmum"`
	Kemandirian           string `json:"kemandirian"`
	Ketangguhan           string `json:"ketangguhan"`
	KetelitianKerja       string `json:"ketelitianKerja"`
	MotivasiBerprestasi   string `json:"motivasiBerprestasi"`
	PenalaranNonVerbal    string `json:"penalaranNonVerbal"`
	PenalaranNumerik      string `json:"penalaranNumerik"`
	PenalaranVerbal       string `json:"penalaranVerbal"`
	PenyesuaianDiri       string `json:"penyesuaianDiri"`
	SistematikaKerja      string `json:"sistematikaKerja"`
}

type CreateIprosTestRequest struct {
	TestID        *int            `json:"testId"`
	Result        IprosTestResult `json:"result" binding:"required"`
	TestTakenDate *time.Time      `json:"testTakenDate"`
	Email         string          `json:"email" binding:"required,email"`
}
