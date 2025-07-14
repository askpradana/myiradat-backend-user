package model

import "time"

type IprosTest struct {
	ID            int        `gorm:"primaryKey" json:"id"`
	TestID        *int       `json:"test_id"`
	Riasec        string     `json:"riasec"`
	Result        []byte     `gorm:"type:jsonb" json:"result"`
	TestTakenDate *time.Time `json:"test_taken_date"`
	ProfileID     *int       `gorm:"unique" json:"profile_id"`
	IsDeleted     bool       `json:"is_deleted"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	CreatedBy     string     `json:"created_by"`
	ModifiedAt    time.Time  `gorm:"autoUpdateTime"`
	ModifiedBy    string     `json:"modified_by"`

	Profile *Profile `gorm:"foreignKey:ProfileID" json:"-"`
}
