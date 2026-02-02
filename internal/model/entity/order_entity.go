package entity

import (
	"gorm.io/plugin/soft_delete"
)

type Order struct {
	BaseEntity
	ProductName string                `json:"product_name"`
	Price       string                `json:"price"`
	UserID      uint                  `json:"user_id"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}
