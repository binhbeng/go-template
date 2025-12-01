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
	model.NewOrderRepository,
)

var serviceSet= wire.NewSet(
	service.NewUserService,
	service.NewOrderService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewOrderHandler,
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

func NewWireGrpc() (*service.GrpcDeps, error) {
    wire.Build(
        dataSet,
        repositorieSet,
        serviceSet,
        wire.Struct(new(service.GrpcDeps), "*"),
    )
    return nil, nil
}