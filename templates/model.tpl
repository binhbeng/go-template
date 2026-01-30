package model

import (
	"gorm.io/plugin/soft_delete"
)

type {{.Pascal}} struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
}

type {{.Pascal}}Repository struct {
	*Repository
}

func New{{.Pascal}}Repository(r *Repository) *{{.Pascal}}Repository {
	return &{{.Pascal}}Repository{
		Repository: r,
	}
}

func (m *{{.Pascal}}Repository) TableName() string {
	return "{{.Name}}s"
}