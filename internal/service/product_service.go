package service

import (
	"github.com/binhbeng/goex/internal/model/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(
	productRepo *repository.ProductRepository,
) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}
