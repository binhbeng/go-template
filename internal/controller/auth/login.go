package auth

import (
	"github.com/binhbeng/goex/internal/controller"
	"github.com/binhbeng/goex/internal/service/auth"
	"github.com/binhbeng/goex/internal/validator"
	"github.com/binhbeng/goex/internal/validator/form"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	controller.Api
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (api LoginController) Login(c *gin.Context) {
	loginForm := form.NewLoginForm()
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return
	}
	result, err := auth.NewLoginService().Login(loginForm.UserName, loginForm.PassWord)
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c, result)
}
