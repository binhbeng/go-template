package app

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/db/sqlc"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/repository"
	"github.com/binhbeng/goex/internal/service"
)

type UserModule struct {
	userHandler *handler.UserHandler
}

func NewUserModule() *UserModule {
	sqlRepo := repository.NewUserRepository(sqlc.New(data.PgxDB))
	userService := service.NewUserService(sqlRepo, data.RedisDB)
	userHandler := handler.NewUserHandler(userService)
	return &UserModule{userHandler: userHandler}
}

func (m *UserModule) Handler() *handler.UserHandler {
	return m.userHandler
}
