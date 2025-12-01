package service

import (
	"context"
	"fmt"

	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/proto"
)

type OrderService interface {
	SayOrder(name string) string
}

type orderService struct {
	proto.UnimplementedOrderServiceServer
	orderRepo *model.OrderRepository
}

func NewOrderService(orderRepo *model.OrderRepository) *orderService {
    return &orderService{
        orderRepo: orderRepo,
    }
}

func (s *orderService) SayOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderReply, error) {
    return &proto.OrderReply{
        Message: fmt.Sprintf("Hello order %v", req.Name),
    }, nil
}
