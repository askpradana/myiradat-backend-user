package model

import "time"

type Consult struct {
	ID             int        `gorm:"primaryKey" json:"id"`
	Owner          string     `json:"owner"`
	ConsultDate    *time.Time `json:"consult_date"`
	AnalysisResult string     `json:"analysis_result"`
	ProfileID      *int       `json:"profile_id"`
	IsDeleted      bool       `json:"is_deleted"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	CreatedBy      string     `json:"created_by"`
	ModifiedAt     time.Time  `gorm:"autoUpdateTime"`
	ModifiedBy     string     `json:"modified_by"`

	Profile *Profile `gorm:"foreignKey:ProfileID" json:"-"`
}
