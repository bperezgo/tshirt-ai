package repositories

import (
	"context"

	"github.com/bperezgo/tshirt_ai/internal/domain/entities"
)

type InMemoryCustomizedProductsRepository struct {
	customizedProducts []*entities.CustomizedProduct
}

func NewCustomizedInMemoryProductsRepository() *InMemoryCustomizedProductsRepository {
	return &InMemoryCustomizedProductsRepository{
		customizedProducts: customized_products_data,
	}
}

func (r *InMemoryCustomizedProductsRepository) Create(ctx context.Context, product entities.CustomizedProduct) error {
	r.customizedProducts = append(r.customizedProducts, &product)
	return nil
}
