package entity

import (
	"gorm.io/plugin/soft_delete"
)

type User struct {
	BaseEntity
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
	Username  string                `json:"username"`
	Password  string                `json:"-"`
	Email     string                `json:"email"`
}
