package queries

import (
	"context"
	"customer-subgraph/graph/model"
)

// Customer resolves the query for a single Customer by ID.
func Customer(ctx context.Context, id string) (*model.Customer, error) {
	return &model.Customer{
		ID: id,
	}, nil
}
