package ports

import (
	"context"

	"github.com/bperezgo/tshirt_ai/internal/domain/entities"
)

type FindByCriteriaProductsRepository struct {
	ID string
}

type ProductsRepository interface {
	FindAll(ctx context.Context) ([]*entities.Product, error)

	FindByCriteria(ctx context.Context, criteria FindByCriteriaProductsRepository) (*entities.Product, error)
}

type CustomizedProductsRepository interface {
	Create(ctx context.Context, product entities.CustomizedProduct) error
}
