package dto

type ProfileServiceRoleInput struct {
	ServiceID int `json:"serviceId" binding:"required"`
	RoleID    int `json:"roleId" binding:"required"`
}

type CreateProfileRequest struct {
	Name     string                    `json:"name" binding:"required"`
	Email    string                    `json:"email" binding:"required,email"`
	NoHP     string                    `json:"noHp" binding:"required"`
	Password string                    `json:"password" binding:"required"`
	Services []ProfileServiceRoleInput `json:"services" binding:"required,dive"`
}
