package form

import "github.com/binhbeng/goex/internal/model"

type LoginAuth struct {
	Username string `form:"username" json:"username"  binding:"required,min=5"` 
	Password string `form:"password" json:"password"  binding:"required,min=6"`
}

type LoginResponse struct {
	User        model.User `json:"user"`
	AccessToken string     `json:"access_token"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}