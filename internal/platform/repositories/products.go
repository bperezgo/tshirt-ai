package repositories

import "github.com/bperezgo/tshirt_ai/internal/domain"

type InMemoryProductsRepository struct {
	ProductsData []*domain.Product
}

func NewInMemoryProductsRepository() *InMemoryProductsRepository {
	return &InMemoryProductsRepository{
		ProductsData: products_data,
	}
}

func (r *InMemoryProductsRepository) FindAll() ([]*domain.Product, error) {
	return r.ProductsData, nil
}
