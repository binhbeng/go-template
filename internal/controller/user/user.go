package user

import (
	"github.com/binhbeng/goex/internal/controller"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/errors"
	"github.com/binhbeng/goex/internal/service/user"
	"github.com/binhbeng/goex/internal/validator"
	"github.com/binhbeng/goex/internal/validator/form"

	// "github.com/binhbeng/goex/internal/service/auth"
	// "github.com/binhbeng/goex/internal/validator"
	// "github.com/binhbeng/goex/internal/validator/form"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller.Api
}

func NewUserController() *UserController {
	return &UserController{}
}

func (api UserController) Me(c *gin.Context) {
	userId := c.GetUint("user_id")
	userModel := model.NewUser()

	user := userModel.GetUserById(userId)

	if err := user.DB().Where("id = ?", userId).First(&user).Error; err != nil {
		api.Err(c, errors.NewBusinessError(errors.FAILURE, err.Error()))
		return
	}
	api.Success(c, user)
}

func (api UserController) Update(c *gin.Context) {
	userId := c.GetUint("user_id")

	updateUserForm := form.NewUpdateUserForm()
	if err := validator.CheckPostParams(c, &updateUserForm); err != nil {
		return
	}

	err := user.NewUserService().Update(userId, updateUserForm)
	
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, nil)
}