package model

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"gorm.io/plugin/soft_delete"
)

type Order struct {
	BaseModel
	ProductName string                `json:"product_name"`
	Price       string                `json:"price"`
	UserID      uint                  `json:"user_id"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}

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
