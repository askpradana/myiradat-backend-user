package model

import "time"

type Service struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	ServiceName string     `json:"service_name"`
	RedirectURI string     `json:"redirect_uri"`
	IsDeleted   bool       `json:"is_deleted"`
	CreatedAt   *time.Time `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	ModifiedAt  *time.Time `json:"modified_at"`
	ModifiedBy  string     `json:"modified_by"`

	Roles []Role `gorm:"foreignKey:MasterServiceID" json:"-"`
}
