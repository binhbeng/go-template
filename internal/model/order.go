package model

type Order struct {
	BaseModelWithSoftDelete
	Productname string `json:"product_name"`
	UserId      uint   `json:"user_id"`
	Price    string `json:"price"`
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

func (m *OrderRepository) GetOrderById(id uint) (*Order, error) {
	var order Order
	if err := m.DB().Model(&Order{}).First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
