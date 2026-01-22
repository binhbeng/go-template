package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/utils/token"
	"github.com/go-redis/redis/v8"
)

// type UserService interface {
// 	Login(ctx context.Context, username, password string) (model.User, string, error)
// 	Me(ctx context.Context, userId int64) (model.User, error)
// 	UpdateProfile(ctx context.Context, userId int64, input dto.UpdateUserInput) (model.User, error)
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

func (s *UserService) Login(ctx context.Context, username, password string) (model.User, string, error) {
	ctx, cancel := context.WithTimeout(ctx, global.DefaultRequestTimeout)
	defer cancel()

	var user model.User
	if err := s.userRepo.DB().WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
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

func (s *UserService) Me(ctx context.Context, userId int64) (model.User, error) {
	user, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userId int64, input dto.UpdateUserInput) (model.User, error) {
	user, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, err
	}

	if err := s.userRepo.DB(&user).WithContext(ctx).Updates(input).Error; err != nil {
		return model.User{}, err
	}

	go func() {
		err := s.redis.Set(ctx, fmt.Sprintf("last_updated:%d", userId), time.Now().UTC(), 0).Err()
		if err != nil {
			log.Println(err)
		}
	}()

	return user, nil
}

func (s *UserService) GetListUser(ctx context.Context, req dto.QueryUsersInput) ([]model.User, error) {
	data, err := s.userRepo.GetListUser(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return data, nil
}
