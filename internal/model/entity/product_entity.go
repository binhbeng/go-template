package entity

import (
	"gorm.io/plugin/soft_delete"
	"github.com/binhbeng/goex/internal/utils/timeutil"
)

type Product struct {
	ID        int                   `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	Name      string                `json:"name"`
	Price     string                `json:"price"`
	CreatedAt timeutil.FormatDate   `gorm:"column:created_at;type:timestamp;<-:create" json:"created_at"`
	UpdatedAt timeutil.FormatDate   `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}
