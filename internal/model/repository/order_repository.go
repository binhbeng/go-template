package repository

import (
	"context"

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

func (m *OrderRepository) GetListOrder(ctx context.Context, filter dto.QueryOrdersInput) ([]dto.GetListOrderResponse, error) {
	var data []dto.GetListOrderResponse

	query := m.DB().WithContext(ctx).Table("orders AS o").
		Select("o.id, o.product_name, o.price, u.id as user_id, u.username, u.email").
		Joins("INNER JOIN users u ON o.user_id = u.id")

	err := query.Scopes(m.Paginate(filter.PageOptionsDto)).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
