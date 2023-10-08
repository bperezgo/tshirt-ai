package repositories

import (
	"context"

	"github.com/bperezgo/tshirt_ai/internal/domain"
)

type InMemoryCustomizedProductsRepository struct {
	customizedProducts []*domain.CustomizedProduct
}

func NewCustomizedInMemoryProductsRepository() *InMemoryCustomizedProductsRepository {
	return &InMemoryCustomizedProductsRepository{
		customizedProducts: customized_products_data,
	}
}

func (r *InMemoryCustomizedProductsRepository) Create(ctx context.Context, product domain.CustomizedProduct) error {
	r.customizedProducts = append(r.customizedProducts, &product)
	return nil
}
