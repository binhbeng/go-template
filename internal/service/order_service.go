package service

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/entity"
	"github.com/binhbeng/goex/internal/model/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
	userRepo  *repository.UserRepository
}

func NewOrderService(
	orderRepo *repository.OrderRepository,
	userRepo *repository.UserRepository,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (s *OrderService) GetListOrders(ctx context.Context, req dto.QueryOrdersInput) ([]dto.GetListOrderResponse, int, error) {
	data, total, err := s.orderRepo.GetOrders(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, userID int, input dto.CreateOrderInput) (entity.Order, error) {
	var createdOrder entity.Order

	err := repository.DB().Transaction(func(tx *gorm.DB) error {
		orderRepoTx := s.orderRepo.WithTx(tx)
		userRepoTx := s.userRepo.WithTx(tx)

		order := entity.Order{
			UserID:    userID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
			Price:     input.Price,
		}

		order, err := orderRepoTx.CreateOrder(ctx, order)
		if err != nil {
			return err
		}

		createdOrder = order
		totalAmount := decimal.RequireFromString(input.Price).Mul(decimal.NewFromInt(int64(input.Quantity))).String()

		_, err = userRepoTx.AddBalance(ctx, userID, totalAmount)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Order{}, err
	}

	return createdOrder, nil
}
