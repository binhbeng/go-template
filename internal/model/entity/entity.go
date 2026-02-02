package entity

import (
	"github.com/binhbeng/goex/internal/utils"
	"gorm.io/plugin/soft_delete"
)

type BaseEntity struct {
	ID        int64            `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreatedAt utils.FormatDate `gorm:"column:created_at;type:timestamp;<-:create" json:"created_at"`
	UpdatedAt utils.FormatDate `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

type BaseEntityWithSoftDelete struct {
	BaseEntity
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}
