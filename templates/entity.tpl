package entity

import (
	"gorm.io/plugin/soft_delete"
)

type {{.Pascal}} struct {
	BaseEntity
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}
