package app

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/model/repository"
	"github.com/binhbeng/goex/internal/service"
)

type UserModule struct {
	userHandler *handler.UserHandler
}

func NewUserModule() *UserModule {
	userRepository := repository.NewUserRepository(data.PostgreDB)
	userService := service.NewUserService(userRepository, data.RedisDB)
	userHandler := handler.NewUserHandler(userService)
	return &UserModule{userHandler: userHandler}
}

func (m *UserModule) Handler() *handler.UserHandler {
	return m.userHandler
}
