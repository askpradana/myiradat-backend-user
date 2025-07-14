package dto

type RoleDTO struct {
	RoleID   int    `json:"roleId"`
	RoleName string `json:"roleName"`
}

type ServiceWithRolesDTO struct {
	ServiceID   int       `json:"serviceId"`
	ServiceName string    `json:"serviceName"`
	Roles       []RoleDTO `json:"roles"`
}
