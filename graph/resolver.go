package graph

import "github.com/bperezgo/tshirt_ai/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	products []*model.Product
}

func NewResolver() *Resolver {
	return &Resolver{
		products: []*model.Product{
			{
				ID:    "1",
				Title: "T-Shirt",
				Price: 10.0,
			},
		},
	}
}
