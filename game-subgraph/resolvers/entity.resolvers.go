package resolvers

import (
	"context"
	"game-subgraph/graph"
	"game-subgraph/graph/model"
	"game-subgraph/resolvers/entities"
)

// FindGameByID is the resolver for the findGameByID field.
func (r *entityResolver) FindGameByID(ctx context.Context, id string) (*model.Game, error) {
	return entities.FindGameByID(ctx, r.GameAPI, id)
}

// Entity returns graph.EntityResolver implementation.
func (r *Resolver) Entity() graph.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
