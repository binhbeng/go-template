package handler

import (
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/internal/utils"
	"github.com/binhbeng/goex/internal/validation"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	loginForm := new(dto.LoginInput)
	if err := validation.ValidateBodyParams(c, &loginForm); err != nil {
		return
	}

	user, accessToken, err := h.userService.Login(ctx, loginForm.Username, loginForm.Password)
	if err != nil {
		utils.HttpBadRequest(c, "Login failed", err)
		return
	}

	response := dto.LoginResponse{
		User:        dto.UserResponse{Id: user.ID, Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt},
		AccessToken: accessToken,
	}

	utils.SuccessResponse(c, 200, "OK", response)
}

// @Summary Get Profile
// @Security BearerAuth
// @Description Returns User
// @Tags User
// @Produce json
// @Success 200 {object} dto.UserResponse
// @Router /user/me [get]
func (h *UserHandler) Me(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.GetInt64("user_id")
	user, err := h.userService.Me(ctx, userId)

	if err != nil {
		utils.HttpBadRequest(c, "Get profile failed", err)
		return
	}

	utils.SuccessResponse(c, 200, "OK", user)
}

// @Summary Update Profile
// @Security BearerAuth
// @Description Returns Status of Update
// @Tags User
// @Produce plain
// @Success 200 {string} string
// @Router /user [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.GetInt64("user_id")
	var updateUserForm dto.UpdateUserInput
	if err := validation.ValidateBodyParams(c, &updateUserForm); err != nil {
		return
	}

	user, err := h.userService.UpdateProfile(ctx, userId, updateUserForm)
	if err != nil {
		utils.HttpBadRequest(c, "Update failed", err)
		return
	}

	utils.SuccessResponse(c, 200, "OK", user)
}

func (h *UserHandler) GetListUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.QueryUsersInput
	if err := validation.ValidateQueryParams(c, &req); err != nil {
		return
	}

	user, err := h.userService.GetListUser(ctx, req)
	if err != nil {
		utils.HttpBadRequest(c, "get failed", err)
		return
	}

	utils.SuccessResponse(c, 200, "OK", user)
}
