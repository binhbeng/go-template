package app

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/model/repository"
	"github.com/binhbeng/goex/internal/service"
)

type ProductModule struct {
	productHandler *handler.ProductHandler
}

func NewProductModule() *ProductModule {
	productRepository := repository.NewProductRepository(data.PostgreDB)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	return &ProductModule{
        productHandler: productHandler,
    }
}

func (m *ProductModule) Handler() *handler.ProductHandler {
	return m.productHandler
}