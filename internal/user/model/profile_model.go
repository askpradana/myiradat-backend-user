package model

import "time"

type Profile struct {
	ID           int        `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	NoHP         string     `json:"no_hp"`
	Password     string     `json:"password"`
	RefreshToken string     `json:"refresh_token"`
	IsDeleted    bool       `json:"is_deleted"`
	CreatedAt    *time.Time `json:"created_at"`
	CreatedBy    string     `json:"created_by"`
	ModifiedAt   *time.Time `json:"modified_at"`
	ModifiedBy   string     `json:"modified_by"`
}
