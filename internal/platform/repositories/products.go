package repositories

import (
	"context"

	"github.com/bperezgo/tshirt_ai/internal/domain/entities"
	"github.com/bperezgo/tshirt_ai/internal/ports"
)

type InMemoryProductsRepository struct {
	ProductsData []*entities.Product
}

func NewInMemoryProductsRepository() *InMemoryProductsRepository {
	return &InMemoryProductsRepository{
		ProductsData: products_data,
	}
}

func (r *InMemoryProductsRepository) FindAll(ctx context.Context) ([]*entities.Product, error) {
	return r.ProductsData, nil
}

func (r *InMemoryProductsRepository) FindByCriteria(ctx context.Context, criteria ports.FindByCriteriaProductsRepository) (*entities.Product, error) {
	for _, product := range r.ProductsData {
		if product.ID == criteria.ID {
			return product, nil
		}
	}

	return nil, nil
}
