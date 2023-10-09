package services

import (
	"context"

	"github.com/bperezgo/tshirt_ai/internal/domain/entities"
	"github.com/bperezgo/tshirt_ai/internal/ports"
)

type CustomizedProductService struct {
	productRepo ports.CustomizedProductsRepository
}

func NewCustomizedProductService(productRepo ports.CustomizedProductsRepository) CustomizedProductService {
	return CustomizedProductService{productRepo: productRepo}
}

func (s *CustomizedProductService) Create(ctx context.Context, createDto entities.CustomizedProduct) error {
	return s.productRepo.Create(ctx, createDto)
}
