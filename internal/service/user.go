package service

import (
	"fmt"
	"time"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// type UserService interface {
// 	Login(username, password string) (*form.LoginResponse, error)
// 	Me(c *gin.Context, userId uint) (*form.UserResponse, error)
// 	UpdateProfile(c *gin.Context, userId uint, data *form.UpdateUserRequest) error
// }

type UserService struct {
	userRepo *model.UserRepository
	redis    *redis.Client
}

func NewUserService(
	userRepo *model.UserRepository,
	redis *redis.Client,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		redis:    redis,
	}
}

func (s *UserService) Login(username, password string) (model.User, string, error) {
	var user model.User
	if err := s.userRepo.DB().Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, "", err
	}

	now := time.Now()
	expiresAt := now.Add(24 * 30 * time.Hour)
	claim := token.NewCustomClaims(&user, expiresAt)
	accessToken, err := token.Generate(claim)

	if err != nil {
		return model.User{}, "", err
	}

	return user, accessToken, nil
}

func (s *UserService) Me(c *gin.Context, userId int64) (model.User, error) {
	user, err := s.userRepo.GetUserById(userId)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *UserService) UpdateProfile(c *gin.Context, userId int64, input dto.UpdateUserInput) (model.User, error) {
	now := time.Now()
	s.redis.Set(c, "last_updated:"+fmt.Sprint(userId), now, 0)
	user, err := s.userRepo.GetUserById(userId)

	if err != nil {
		return model.User{}, err
	}
	
	db := s.userRepo.DB(&model.User{})

	db.Model(&user).Updates(input)

	if err := db.Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
