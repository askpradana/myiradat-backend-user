package model

import "time"

type Comment struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Comment    string    `json:"comment"`
	ConsultID  *int      `json:"consult_id"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
	ModifiedBy string    `json:"modified_by"`

	Consult *Consult `gorm:"foreignKey:ConsultID" json:"-"`
}
