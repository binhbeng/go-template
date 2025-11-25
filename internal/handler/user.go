package handler

import (
	"github.com/binhbeng/goex/internal/api"
	"github.com/binhbeng/goex/internal/api/form"
	"github.com/binhbeng/goex/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	loginForm := new(form.LoginAuth)
	if err := api.CheckPostParams(c, &loginForm); err != nil {
		return
	}

	data, err := h.userService.Login(loginForm.Username, loginForm.Password)
	if err != nil {
		api.HandleError(c, 500, "Login failed", err, nil)
		return
	}

	api.HandleSuccess(c, 200, "OK", data)
}

func (h *UserHandler) Me(c *gin.Context) {
	userId := GetUserIdFromCtx(c)
	user, err := h.userService.Me(c, userId)

	if err != nil {
		api.HandleError(c, 500, "Get profile failed", err, nil)
		return
	}

	api.HandleSuccess(c, 200, "OK", user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userId := GetUserIdFromCtx(c)
	var updateUserForm form.UpdateUserRequest
	if err := api.CheckPostParams(c, &updateUserForm); err != nil {
		return
	}

	if err := h.userService.UpdateProfile(c, userId, &updateUserForm); err != nil {
		api.HandleError(c, 500, "Update failed", err, nil)
		return
	}

	api.HandleSuccess(c, 200, "OK")
}
