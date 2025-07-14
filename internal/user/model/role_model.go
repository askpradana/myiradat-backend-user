package model

import "time"

type Role struct {
	ID              int        `gorm:"primaryKey" json:"id"`
	RoleName        string     `json:"role_name"`
	Description     string     `json:"description"`
	MasterServiceID *int       `json:"master_service_id"`
	IsDeleted       bool       `json:"is_deleted"`
	CreatedAt       *time.Time `json:"created_at"`
	CreatedBy       string     `json:"created_by"`
	ModifiedAt      *time.Time `json:"modified_at"`
	ModifiedBy      string     `json:"modified_by"`

	MasterService *Service `gorm:"foreignKey:MasterServiceID" json:"-"`
}
