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

func (s *OrderService) GetListOrder(ctx context.Context, req dto.QueryOrdersInput) (any, error) {
	data, err := s.orderRepo.GetListOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
