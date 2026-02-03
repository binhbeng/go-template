package service

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderService(
	orderRepo *repository.OrderRepository,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) GetListOrders(ctx context.Context, req dto.QueryOrdersInput) ([]dto.GetListOrderResponse, int, error) {
	data, total, err := s.orderRepo.GetOrders(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
