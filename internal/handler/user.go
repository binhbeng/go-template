package handler

import (
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/response"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/internal/validator"
	"github.com/binhbeng/goex/internal/validator/form"
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
	// panic("hehe")
	loginForm := form.NewLoginForm()
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return
	}

	data, err := h.userService.Login(loginForm.Username, loginForm.Password)
	if err != nil {
		response.ErrorResponse(c, 500, "Error hehe", err, nil)
		return
	}

	response.SuccessResponse(c, 200, "OK", data)
}

func (h *UserHandler) Me(c *gin.Context) {
	userId := GetUserIdFromCtx(c)
	h.userService.Me(c, userId)
}

func (h *UserHandler) Update(c *gin.Context) {
	userId := GetUserIdFromCtx(c)
	updateUserForm := form.NewUpdateUserForm()
	if err := validator.CheckPostParams(c, &updateUserForm); err != nil {
		return
	}

	h.userService.Update(c, userId, &model.User{
		Email: updateUserForm.Email,
	})
}
