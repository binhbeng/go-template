//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/internal/router"
	"github.com/google/wire"
)

var repositorieSet = wire.NewSet(
	model.NewRepository,
	model.NewUserRepository,
)

var serviceSet= wire.NewSet(
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var dataSet = wire.NewSet(
	data.NewRedis,
	data.NewPostgreDB,
)

func NewWire() (*router.RouterDeps, error) {
	wire.Build(
		dataSet,
		repositorieSet,
		serviceSet,
		handlerSet,
		wire.Struct(new(router.RouterDeps), "*"),
	)

	return nil, nil
}
