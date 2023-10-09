// This module is temporal, while I switch the in memory repository to the database repository.
package repositories

import "github.com/bperezgo/tshirt_ai/internal/domain/entities"

var products_data = []*entities.Product{
	{
		ID:          "1",
		Title:       "T-Shirt 1",
		Description: "T-Shirt 1 description",
		Price:       10.0,
		Images: []string{
			"https://images.unsplash.com/photo-1593642532452-9d5c0b2b8b0f?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60",
			"https://images.unsplash.com/photo-1593642532452-9d5c0b2b8b0f?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60",
		},
	},
}
