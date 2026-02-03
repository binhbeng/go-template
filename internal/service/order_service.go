package service

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/entity"
	"github.com/binhbeng/goex/internal/model/repository"
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
			ProductName: input.ProductName,
			Price:       input.Price,
			UserID:      userID,
		}

		order, err := orderRepoTx.CreateOrder(ctx, order)
		if err != nil {
			return err
		}

		createdOrder = order

		user, err := userRepoTx.GetUserById(ctx, userID)
		if err != nil {
			return err
		}

		user.Balance += input.Price
		
		if err := userRepoTx.DB.WithContext(ctx).Save(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.Order{}, err
	}

	return createdOrder, nil
}
