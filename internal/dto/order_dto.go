package dto

type OrderInput struct {
	Example string `form:"example" json:"example"`
}

type QueryOrdersInput struct {
	Search string `json:"search"`
	PageOptionsDto
}

type GetListOrderResponse struct {
	ID          int   `json:"id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	UserID      int   `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}

type CreateOrderInput struct {
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
}