package dto

type UpdateProfileWithRolesRequest struct {
	ProfileID int    `json:"profileId" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	NoHP      string `json:"noHp" binding:"required"`
	Password  string `json:"password"` // optional

	Services []struct {
		ServiceID int `json:"serviceId" binding:"required"`
		RoleID    int `json:"roleId" binding:"required"`
	} `json:"services" binding:"required"`
}

type UpdateProfileRequest struct {
	ProfileID int    `json:"profileId" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	NoHP      string `json:"noHp" binding:"required"`
	Password  string `json:"password"` // optional
}
