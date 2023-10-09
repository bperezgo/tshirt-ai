package graph

import (
	"github.com/bperezgo/tshirt_ai/internal/domain/services"
	"github.com/bperezgo/tshirt_ai/internal/platform/repositories"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productsService           services.ProductService
	customizedProductsService services.CustomizedProductService
}

func NewResolver() *Resolver {
	productsRepository := repositories.NewInMemoryProductsRepository()
	customizedProductsRepository := repositories.NewCustomizedInMemoryProductsRepository()

	productsService := services.NewProductService(productsRepository)
	customizedProductsService := services.NewCustomizedProductService(customizedProductsRepository)

	return &Resolver{
		productsService:           productsService,
		customizedProductsService: customizedProductsService,
	}
}
