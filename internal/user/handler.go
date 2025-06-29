package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vfyuliawan/my-golang-sdk/core/api_constant"
	"github.com/vfyuliawan/my-golang-sdk/core/json_response"
	"net/http"
)

func HttpHandler(router *gin.Engine) {
	handler := NewHandler()
	user := router.Group("/user")
	{
		user.GET("get-user-detail", handler.GetUserDetail)
	}

}

var service = func() InterfaceUserService {
	return NewUserService()
}

type Handler struct {
	service InterfaceUserService
}

// GetUserDetail godoc
// @Summary Get user detail
// @Description Get detail of a user by ID or session context
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ResponseGetDetail
// @Failure 400 {object} any
// @Router /user/get-user-detail [get]
func (h *Handler) GetUserDetail(c *gin.Context) {
	res, err := h.service.GetUserDetail(c, RequestGetDetail{})
	if err != nil {
		json_response.Error("User Service", api_constant.InvalidRequest.GetCode(), err.Error(), http.StatusBadRequest)
	}

	json_response.Success(c, res)

}

func NewHandler() *Handler {
	return &Handler{service: service()}
}
