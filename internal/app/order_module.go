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
	baseRepo := repository.NewRepository(data.PostgreDB)
	orderRepository := repository.NewOrderRepository(baseRepo)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)
	return &OrderModule{
        orderHandler: orderHandler,
    }
}

func (m *OrderModule) Handler() *handler.OrderHandler {
	return m.orderHandler
}