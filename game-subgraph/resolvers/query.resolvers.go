package resolvers

import (
	"context"
	"game-subgraph/graph"
	"game-subgraph/graph/model"
	"game-subgraph/resolvers/queries"
)

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Game is the resolver for the game field.
func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	return queries.Game(ctx, r.GameAPI, id)
}
