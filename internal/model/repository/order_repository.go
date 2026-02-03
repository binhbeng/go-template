package repository

import (
	"context"
	"fmt"

	"github.com/binhbeng/goex/internal/dto"
)

type OrderRepository struct {
	*Repository
}

func NewOrderRepository(r *Repository) *OrderRepository {
	return &OrderRepository{
		Repository: r,
	}
}

func (m *OrderRepository) TableName() string {
	return "orders"
}

func (m *OrderRepository) GetOrders(ctx context.Context, filter dto.QueryOrdersInput) ([]dto.GetListOrderResponse, int, error) {
	var orderList []dto.GetListOrderResponse
	var totalCount int64

	query := m.DB().WithContext(ctx).
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

	err = query.Scopes(m.Paginate(filter.PageOptionsDto)).Find(&orderList).Error
	if err != nil {
		return nil, 0, err
	}

	return orderList, int(totalCount), nil
}
