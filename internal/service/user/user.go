package user

import (

	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/errors"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/internal/validator/form"
)

type UserService struct {
	service.Base
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Update(userId uint, userForm *form.UpdateUser) error {
	form := map[string]any{
		"email": userForm.Email,
	}

	if err:= data.PostgreDB.Model(&model.User{}).Where("id = ?", userId).Updates(form).Error; err != nil {
		return errors.NewBusinessError(444)
	}

	return nil
}
