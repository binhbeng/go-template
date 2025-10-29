package auth

import (
	"time"

	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/errors"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
	"github.com/binhbeng/goex/internal/service"
)

type LoginResponse struct {
	User model.User `json:"user"`
	AccessToken string `json:"access_token"`
}

type LoginService struct {
	service.Base
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) Login(username, password string) (*LoginResponse, error) {
	var user model.User
	if err := user.DB().Where("username = ?", username).First(&user).Error; err != nil {
		err := errors.NewBusinessError(errors.UserDoesNotExist)
		return nil, err
	}
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)
	claim := token.NewCustomClaims(&user, expiresAt)
	accessToken, err := token.Generate(claim)

	if err != nil {
		return nil, errors.NewBusinessError(errors.FAILURE, err.Error())
	}

	return &LoginResponse{
		User:        user,
		AccessToken: accessToken,
	}, nil

}
