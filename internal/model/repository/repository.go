package repository

import (
	"github.com/binhbeng/goex/data"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func DB() *gorm.DB {
	return data.PostgreDB
}

func (m *Repository) DB(entity ...any) *gorm.DB {
	if entity != nil {
		return m.db.Model(entity[0])
	}
	return m.db
}

