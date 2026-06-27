package resolvers

import (
	"context"
	"customer-subgraph/graph"
	"customer-subgraph/graph/model"
	"customer-subgraph/resolvers/entities"
)

// Customer returns graph.CustomerResolver implementation.
func (r *Resolver) Customer() graph.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }

// Transactions is the resolver for the transactions field.
func (r *customerResolver) Transactions(ctx context.Context, obj *model.Customer) ([]*model.Transaction, error) {
	return entities.CustomerTransactions(ctx, r.CustomerAPI, obj)
}
