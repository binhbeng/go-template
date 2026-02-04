package dto

type OrderInput struct {
	Example string `form:"example" json:"example"`
}

type QueryOrdersInput struct {
	Search string `json:"search"`
	PageOptionsDto
}

type GetListOrderResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Quantity    int    `json:"quantity"`
}

type CreateOrderInput struct {
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
}
