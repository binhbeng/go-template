package entity

import (
	"github.com/binhbeng/goex/internal/utils/timeutil"
	"gorm.io/plugin/soft_delete"
)

type Order struct {
	ID          int                   `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	ProductName string                `json:"product_name"`
	Price       string                `json:"price"`
	UserID      int                   `json:"user_id"`
	CreatedAt   timeutil.FormatDate   `gorm:"column:created_at;type:timestamp;<-:create" json:"created_at"`
	UpdatedAt   timeutil.FormatDate   `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}
