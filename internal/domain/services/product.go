package services

import (
	"context"

	"github.com/bperezgo/tshirt_ai/internal/domain/entities"
	"github.com/bperezgo/tshirt_ai/internal/ports"
)

type ProductService struct {
	productRepo ports.ProductsRepository
}

func NewProductService(productRepo ports.ProductsRepository) ProductService {
	return ProductService{productRepo: productRepo}
}

func (s *ProductService) FindAll(ctx context.Context) ([]*entities.Product, error) {
	return s.productRepo.FindAll(ctx)
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*entities.Product, error) {
	return s.productRepo.FindByCriteria(ctx, ports.FindByCriteriaProductsRepository{ID: id})
}
