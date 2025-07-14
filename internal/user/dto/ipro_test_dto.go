package dto

import "time"

type IproTestResult struct {
	KecerdasanUmum                string `json:"kecerdasanUmum"`
	DayaAnalisaSintesa            string `json:"dayaAnalisaSintesa"`
	DayaBerpikirAbtrak            string `json:"dayaBerpikirAbtrak"`
	LogikaBerpikir                string `json:"logikaBerpikir"`
	KelincahanPikir               string `json:"kelincahanPikir"`
	Inisiatif                     string `json:"inisiatif"`
	PerencanaanDanPerorganisasian string `json:"perencanaanDanPerorganisasian"`
	SistematikaKerja              string `json:"sistematikaKerja"`
	Fleksibilitas                 string `json:"fleksibilitas"`
	DayaTahanKerjaRutin           string `json:"dayaTahanKerjaRutin"`
	DayaTahanKerjaStress          string `json:"dayaTahanKerjaStress"`
	StabilitasEmosi               string `json:"stabilitasEmosi"`
	PenyesuaianDiri               string `json:"penyesuaianDiri"`
	HubunganInterpersonal         string `json:"hubunganInterpersonal"`
	Kerjasama                     string `json:"kerjasama"`
	KepercayaanDiri               string `json:"kepercayaanDiri"`
	Kepemimpinan                  string `json:"kepemimpinan"`
	PengambilanKeputusan          string `json:"pengambilanKeputusan"`
	MotivasiBerprestasi           string `json:"motivasiBerprestasi"`
	KomitmenTugas                 string `json:"komitmenTugas"`
}

type CreateIproTestRequest struct {
	TestID        *int           `json:"testId"`
	Result        IproTestResult `json:"result" binding:"required"`
	TestTakenDate *time.Time     `json:"testTakenDate"`
	Email         string         `json:"email" binding:"required,email"`
}
