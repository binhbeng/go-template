package service

import (
	"fmt"
	"time"

	"github.com/binhbeng/goex/internal/db/sqlc"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
	"github.com/binhbeng/goex/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type UserService struct {
	userRepo *repository.UserRepository
	redis    *redis.Client
}

func NewUserService(
	userRepo *repository.UserRepository,
	redis *redis.Client,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		redis:    redis,
	}
}

func (s *UserService) Login(username, password string) (sqlc.User, string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return sqlc.User{}, "", err
	}

	now := time.Now()
	expiresAt := now.Add(24 * 30 * time.Hour)
	claim := token.NewCustomClaims(user, expiresAt)
	accessToken, err := token.Generate(claim)

	if err != nil {
		return sqlc.User{}, "", err
	}

	return user, accessToken, nil
}

func (s *UserService) Me(c *gin.Context, userId int64) (sqlc.User, error) {
	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (s *UserService) UpdateProfile(c *gin.Context, userId int64, form sqlc.UpdateUserParams) (sqlc.User, error) {
	now := time.Now()
	s.redis.Set(c, "last_updated:"+fmt.Sprint(userId), now, 0)

	updatedUser, err := s.userRepo.UpdateProfile(c, sqlc.UpdateUserParams{
		ID:    userId,
		Email: form.Email,
	})

	if err != nil {
		return sqlc.User{}, err
	}

	return updatedUser, nil
}
