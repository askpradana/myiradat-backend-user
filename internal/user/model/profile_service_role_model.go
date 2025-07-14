package model

type ProfileServiceRole struct {
	ProfileID int `gorm:"primaryKey" json:"profile_id"`
	ServiceID int `gorm:"primaryKey" json:"service_id"`
	RoleID    int `json:"role_id"`

	Profile *Profile `gorm:"foreignKey:ProfileID" json:"-"`
	Service *Service `gorm:"foreignKey:ServiceID" json:"-"`
	Role    *Role    `gorm:"foreignKey:RoleID,ServiceID;references:ID,MasterServiceID" json:"-"`
}
