package resolvers

import (
	"context"
	"customer-subgraph/graph"
	"customer-subgraph/graph/model"
	"customer-subgraph/resolvers/queries"
)

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Customer is the resolver for the customer field.
func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	return queries.Customer(ctx, r.CustomerAPI, id)
}
