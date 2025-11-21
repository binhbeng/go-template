package service

import (
	"time"

	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/pkg/response"
	"github.com/binhbeng/goex/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(username, password string) (*LoginResponse, error)
	Me(c *gin.Context, userId uint)
	Update(c *gin.Context, userId uint, data *model.User)
}

type userService struct {
	userRepo *model.User
	// *Service
}

func NewUserService(
	// service *Service,
	userRepo *model.User,
) UserService {
	return &userService{
		userRepo: userRepo,
		// Service:  service,
	}
}

type LoginResponse struct {
	User        model.User `json:"user"`
	AccessToken string     `json:"access_token"`
}

type UserResponse struct {
	model.User
}

func (u *userService) Login(username, password string) (*LoginResponse, error) {
	// userModel := model.NewUser()
	user := model.User{}
	if err := u.userRepo.DB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)
	claim := token.NewCustomClaims(&user, expiresAt)
	accessToken, err := token.Generate(claim)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:        user,
		AccessToken: accessToken,
	}, nil
}

func (u *userService) Me(c *gin.Context, userId uint) {
	userModel := model.NewUser()
	user := userModel.GetUserById(userId)

	if user != nil {
		response.SuccessResponse(c, 200, "OK", UserResponse{*user})
	}
}

func (u *userService) Update(c *gin.Context, userId uint, form *model.User) {

	if err := u.userRepo.DB().Where("id = ?", userId).Updates(form).Error; err != nil {
		response.ErrorResponse(c, 500, "Error hehe", err, nil)
		return
	}

	response.SuccessResponse(c, 200, "OK", nil)
}
