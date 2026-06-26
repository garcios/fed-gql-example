package queries

import (
	"context"
	"customer-subgraph/datasources"
	"customer-subgraph/graph/model"
)

// Customer resolves the query for a single Customer by ID.
func Customer(ctx context.Context, api *datasources.MockAPI, id string) (*model.Customer, error) {
	return api.GetCustomer(ctx, id)
}
