package user

import (
	"myiradat-backend-auth/internal/configs"
	"myiradat-backend-auth/internal/response"
	"myiradat-backend-auth/internal/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func HttpHandler(r *gin.Engine) {
	// Ambil koneksi DB dari config
	db := configs.Database.DbUser()

	// Buat dependency: repo → service → handler
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	// Register routes
	r.GET("/", handler.HealthCheck)
	r.POST("/profile/summary", handler.GetProfileSummary)
	r.POST("/profile/detail", handler.GetProfileDetail)
	r.GET("/services-with-roles", handler.GetServicesWithRoles)
	r.POST("/profiles", handler.CreateProfile)
	r.POST("/profiles/list", handler.ListProfiles)
	r.POST("/profiles/update-with-roles", handler.UpdateProfileWithRoles)
	r.POST("/profile/update-basic", handler.UpdateProfile)
	r.POST("/ipro-tests", handler.CreateIproTest)
	r.POST("/ipros-tests", handler.CreateIprosTest)
	r.POST("/iprob-tests", handler.CreateIprobTest)
	r.POST("/improve-care", handler.CreateConsult)
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Profile Service is running in docker!"})
}

func (h *Handler) GetProfileSummary(c *gin.Context) {
	var req dto.GetProfileSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "invalid request: email is required and must be valid")
		return
	}

	data, err := h.service.GetProfileSummary(req.Email)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, data)
}

func (h *Handler) GetProfileDetail(c *gin.Context) {
	var req dto.GetProfileDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	result, err := h.service.GetProfileDetail(req.Email)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *Handler) GetServicesWithRoles(c *gin.Context) {
	data, err := h.service.GetServicesWithRoles()
	if err != nil {
		response.ServerError(c, "Failed to fetch services and roles")
		return
	}
	response.Success(c, data)
}

func (h *Handler) ListProfiles(c *gin.Context) {
	var req dto.ListProfilesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	result, err := h.service.ListProfiles(req)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *Handler) CreateProfile(c *gin.Context) {
	var req dto.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.CreateProfile(req); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Profile created successfully"})
}

func (h *Handler) UpdateProfileWithRoles(c *gin.Context) {
	var req dto.UpdateProfileWithRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.UpdateProfileWithRoles(req); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Profile updated successfully"})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.UpdateProfile(req); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Profile updated successfully"})
}

func (h *Handler) CreateIproTest(c *gin.Context) {
	var req dto.CreateIproTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.CreateIproTest(req); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Ipro test created successfully"})
}

func (h *Handler) CreateIprosTest(c *gin.Context) {
	var req dto.CreateIprosTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.CreateIprosTest(req); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Ipros test created successfully"})
}

func (h *Handler) CreateIprobTest(c *gin.Context) {
	var req dto.CreateIprobTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.CreateIprobTest(req); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Iprob test created successfully"})
}

func (h *Handler) CreateConsult(c *gin.Context) {
	var req dto.CreateConsultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err.Error())
		return
	}

	if err := h.service.CreateConsultWithComments(req); err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "consult created successfully"})
}
