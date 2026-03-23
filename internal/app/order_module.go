package app

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/model/repository"
	"github.com/binhbeng/goex/internal/service"
	"github.com/binhbeng/goex/pkg/kafka"
)

type OrderModule struct {
	orderHandler *handler.OrderHandler
	orderService *service.OrderService
}

func NewOrderModule(kafkaProducer kafka.KafkaProducer) *OrderModule {
	orderRepository := repository.NewOrderRepository(data.PostgreDB)
	userRepository := repository.NewUserRepository(data.PostgreDB)

	orderService := service.NewOrderService(orderRepository, userRepository)
	orderHandler := handler.NewOrderHandler(orderService, kafkaProducer)

	return &OrderModule{
		orderHandler: orderHandler,
		orderService: orderService,
	}
}

func (o *OrderModule) Handler() *handler.OrderHandler {
	return o.orderHandler
}

func (o *OrderModule) Service() *service.OrderService {
	return o.orderService
}
