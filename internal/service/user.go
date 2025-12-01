package service

import (
	"fmt"
	"time"

	"github.com/binhbeng/goex/internal/api/form"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type UserService interface {
	Login(username, password string) (*form.LoginResponse, error)
	Me(c *gin.Context, userId uint) (*form.UserResponse, error)
	UpdateProfile(c *gin.Context, userId uint, data *form.UpdateUserRequest) error
}

type userService struct {
	userRepo *model.UserRepository
	redis    *redis.Client
}

func NewUserService(userRepo *model.UserRepository, redis *redis.Client) UserService {
	return &userService{
		userRepo: userRepo,
		redis:    redis,
	}
}

func (s *userService) Login(username, password string) (*form.LoginResponse, error) {
	var user model.User
	if err := s.userRepo.DB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	now := time.Now()
	expiresAt := now.Add(24 * 30 * time.Hour)
	claim := token.NewCustomClaims(&user, expiresAt)
	accessToken, err := token.Generate(claim)

	if err != nil {
		return nil, err
	}

	return &form.LoginResponse{
		User:        user,
		AccessToken: accessToken,
	}, nil
}

func (s *userService) Me(c *gin.Context, userId uint) (*form.UserResponse, error) {
	user, err := s.userRepo.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return &form.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) UpdateProfile(c *gin.Context, userId uint, form *form.UpdateUserRequest) error {
	now := time.Now()
	s.redis.Set(c, "last_updated:"+fmt.Sprint(userId), now, 0)

	return s.userRepo.DB(&model.User{}).Where("id = ?", userId).Updates(form).Error
}
