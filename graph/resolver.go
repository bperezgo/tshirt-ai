package graph

import (
	"github.com/bperezgo/tshirt_ai/internal/platform/repositories"
	"github.com/bperezgo/tshirt_ai/internal/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productsRepository ports.ProductsRepository
}

func NewResolver() *Resolver {
	productsRepository := repositories.NewInMemoryProductsRepository()
	// customizedProductsRepository := repository.NewCustomizedProductsRepository()

	return &Resolver{
		productsRepository: productsRepository,
	}
}
