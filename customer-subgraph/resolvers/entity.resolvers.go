package resolvers

import (
	"context"
	"customer-subgraph/graph"
	"customer-subgraph/graph/model"
	"customer-subgraph/resolvers/entities"
)

// FindCustomerByID is the resolver for the findCustomerByID field.
func (r *entityResolver) FindCustomerByID(ctx context.Context, id string) (*model.Customer, error) {
	return entities.FindCustomerByID(ctx, r.CustomerAPI, id)
}

// FindGameByID is the resolver for the findGameByID field.
func (r *entityResolver) FindGameByID(ctx context.Context, id string) (*model.Game, error) {
	return entities.FindGameByID(ctx, r.CustomerAPI, id)
}

// Entity returns graph.EntityResolver implementation.
func (r *Resolver) Entity() graph.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
