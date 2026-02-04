package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/binhbeng/goex/config"
	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/global"
	"github.com/binhbeng/goex/internal/model/entity"
	"github.com/binhbeng/goex/internal/model/repository"
	"github.com/binhbeng/goex/internal/utils/token"
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

func (s *UserService) Login(ctx context.Context, username, password string) (entity.User, string, error) {
	ctx, cancel := context.WithTimeout(ctx, global.DefaultRequestTimeout)
	defer cancel()

	var user entity.User
	if err := s.userRepo.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, "", err
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(config.Cfg.Jwt.TTL) * time.Second)
	claim := token.NewCustomClaims(&user, expiresAt)
	accessToken, err := token.Generate(claim)

	if err != nil {
		return entity.User{}, "", err
	}

	return user, accessToken, nil
}

func (s *UserService) Me(ctx context.Context, userId int) (entity.User, error) {
	user, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userId int, input dto.UpdateUserInput) (entity.User, error) {
	user, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return entity.User{}, err
	}

	if err := s.userRepo.DB.WithContext(ctx).Updates(input).Error; err != nil {
		return entity.User{}, err
	}

	go func() {
		err := s.redis.Set(ctx, fmt.Sprintf("last_updated:%d", userId), time.Now().UTC(), 0).Err()
		if err != nil {
			log.Println(err)
		}
	}()

	return user, nil
}

func (s *UserService) GetListUser(ctx context.Context, req dto.QueryUsersInput) ([]entity.User, error) {
	data, err := s.userRepo.GetListUser(ctx, req)
	if err != nil {
		return nil, err
	}
	
	return data, nil
}
