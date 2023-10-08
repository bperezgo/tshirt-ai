package ports

import "github.com/bperezgo/tshirt_ai/internal/domain"

type ProductsRepository interface {
	FindAll() ([]*domain.Product, error)
}
