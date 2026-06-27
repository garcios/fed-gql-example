package resolvers

import (
	"context"
	"customer-subgraph/graph"
	"customer-subgraph/graph/model"
	"customer-subgraph/resolvers/entities"
)

// Bet returns graph.BetResolver implementation.
func (r *Resolver) Bet() graph.BetResolver { return &betResolver{r} }

type betResolver struct{ *Resolver }

// Game is the resolver for the game field.
func (r *betResolver) Game(ctx context.Context, obj *model.Bet) (*model.Game, error) {
	return entities.BetGame(ctx, r.CustomerAPI, obj)
}

// Amount is the resolver for the amount field.
func (r *betResolver) Amount(ctx context.Context, obj *model.Bet) (float64, error) {
	return obj.Amount, nil
}
