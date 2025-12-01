package handler

import (
	"context"

	"github.com/binhbeng/goex/internal/proto"
	"github.com/binhbeng/goex/internal/service"
)

type OrderHandler struct {
	*Handler
	proto.OrderServiceServer
	orderService service.OrderService
}

func NewOrderHandler(handler *Handler, orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		Handler:     handler,
		orderService: orderService,
	}
}

func (s *OrderHandler) SayOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderReply, error) {
	msg := s.orderService.SayOrder(req.Name)
	return &proto.OrderReply{Message: msg}, nil
}