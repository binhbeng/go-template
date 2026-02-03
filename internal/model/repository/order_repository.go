package repository

import (
	"context"
	"fmt"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/entity"
	"github.com/binhbeng/goex/internal/utils"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (m *OrderRepository) WithTx(tx *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: tx,
	}
}

func (m *OrderRepository) TableName() string {
	return "orders"
}

func (m *OrderRepository) GetOrders(ctx context.Context, filter dto.QueryOrdersInput) ([]dto.GetListOrderResponse, int, error) {
	var orderList []dto.GetListOrderResponse
	var totalCount int64

	query := m.DB.WithContext(ctx).
		Table("orders AS o").
		Select("o.id, o.product_name, o.price, u.id as user_id, u.username, u.email").
		Joins("INNER JOIN users u ON o.user_id = u.id")

	allowedOrderByColumns := map[string]bool{
		"id":         true,
		"user_id":    true,
		"created_at": true,
	}

	if filter.OrderBy != "" && !allowedOrderByColumns[filter.OrderBy] {
		return nil, 0, fmt.Errorf("invalid order by column: %s", filter.OrderBy)
	}

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Scopes(utils.Paginate(filter.PageOptionsDto)).Find(&orderList).Error
	if err != nil {
		return nil, 0, err
	}

	return orderList, int(totalCount), nil
}

func (m *OrderRepository) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	if err := m.DB.WithContext(ctx).Create(&order).Error; err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
