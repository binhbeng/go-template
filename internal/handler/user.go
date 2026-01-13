package handler

import (
	"github.com/binhbeng/goex/internal/db/sqlc"
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/pkg/utils"
	"github.com/binhbeng/goex/internal/pkg/utils/api"
	"github.com/binhbeng/goex/internal/service"
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
	loginForm := new(dto.LoginInput)
	if err := api.CheckPostParams(c, &loginForm); err != nil {
		return
	}

	user, accessToken, err := h.userService.Login(loginForm.Username, loginForm.Password)
	if err != nil {
		api.HandleError(c, 500, "Login failed", err)
		return
	}

	response := dto.LoginResponse{
		User:        user,
		AccessToken: accessToken,
	}

	api.HandleSuccess(c, 200, "OK", response)
}

// @Summary Get Profile
// @Security BearerAuth
// @Description Returns User
// @Tags User
// @Produce json
// @Success 200 {object} form.UserResponse
// @Router /user/me [get]
func (h *UserHandler) Me(c *gin.Context) {
	userId := api.GetUserIdFromCtx(c)
	user, err := h.userService.Me(c, userId)

	if err != nil {
		api.HandleError(c, 500, "Get profile failed", err)
		return
	}

	response := &dto.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: utils.FormatDate{Time: user.CreatedAt.Time},
	}

	api.HandleSuccess(c, 200, "OK", response)
}

// @Summary Update Profile
// @Security BearerAuth
// @Description Returns Status of Update
// @Tags User
// @Produce plain
// @Success 200 {string} string
// @Router /user [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userId := api.GetUserIdFromCtx(c)
	var updateUserForm dto.UpdateUserInput
	if err := api.CheckPostParams(c, &updateUserForm); err != nil {
		return
	}

	if _, err := h.userService.UpdateProfile(c, userId, sqlc.UpdateUserParams{Email: updateUserForm.Email}); err != nil {
		api.HandleError(c, 500, "Update failed", err)
		return
	}

	api.HandleSuccess(c, 200, "OK")
}
