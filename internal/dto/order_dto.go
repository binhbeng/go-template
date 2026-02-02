package dto

import ()

type OrderInput struct {
	Example string `form:"example" json:"example"`
}

type QueryOrdersInput struct {
	Search string `json:"search"`
	PageOptionsDto
}

type GetListOrderResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}
