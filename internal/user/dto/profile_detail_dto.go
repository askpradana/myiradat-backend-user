package dto

type GetProfileDetailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ServiceWithRole struct {
	ServiceID   int    `json:"serviceId"`
	ServiceName string `json:"serviceName"`
	ServiceCode string `json:"serviceCode"`
	RoleID      int    `json:"roleId"`
	RoleName    string `json:"roleName"`
}

type GetProfileDetailResponse struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	NoHP     string            `json:"noHp"`
	Services []ServiceWithRole `json:"services"`
}
