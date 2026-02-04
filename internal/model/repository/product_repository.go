package repository

import (
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (m *ProductRepository) WithTx(tx *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: tx,
	}
}

func (m *ProductRepository) TableName() string {
	return "products"
}
