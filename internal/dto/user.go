package dto

import (
	"github.com/binhbeng/goex/internal/pkg/utils"
)

type LoginInput struct {
	Username string `form:"username" json:"username"  binding:"required,min=5"`
	Password string `form:"password" json:"password"  binding:"required,min=6"`
}

type UserResponse struct {
	Id        int64            `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt utils.FormatDate `json:"created_at"`
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string     `json:"access_token"`
}

type UpdateUserInput struct {
	Email string `form:"email" json:"email" binding:"required"`
}
