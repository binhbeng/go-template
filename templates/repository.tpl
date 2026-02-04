package repository

import (
	"gorm.io/gorm"
)

type {{.Pascal}}Repository struct {
	DB *gorm.DB
}

func New{{.Pascal}}Repository(db *gorm.DB) *{{.Pascal}}Repository {
	return &{{.Pascal}}Repository{
		DB: db,
	}
}

func (m *{{.Pascal}}Repository) WithTx(tx *gorm.DB) *{{.Pascal}}Repository {
	return &{{.Pascal}}Repository{
		DB: tx,
	}
}

func (m *{{.Pascal}}Repository) TableName() string {
	return "{{.Name}}s"
}