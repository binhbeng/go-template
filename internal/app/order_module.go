package app

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/model/repository"
	"github.com/binhbeng/goex/internal/service"
)

type OrderModule struct {
	orderHandler *handler.OrderHandler
}

func NewOrderModule() *OrderModule {
	orderRepository := repository.NewOrderRepository(data.PostgreDB)
	userRepository := repository.NewUserRepository(data.PostgreDB)
	orderService := service.NewOrderService(orderRepository, userRepository)
	orderHandler := handler.NewOrderHandler(orderService)
	return &OrderModule{
		orderHandler: orderHandler,
	}
}

func (m *OrderModule) Handler() *handler.OrderHandler {
	return m.orderHandler
}
