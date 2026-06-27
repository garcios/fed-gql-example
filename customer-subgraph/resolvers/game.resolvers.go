package resolvers

import (
	"context"
	"customer-subgraph/graph"
	"customer-subgraph/graph/model"
	"customer-subgraph/resolvers/entities"
)

// Game returns graph.GameResolver implementation.
func (r *Resolver) Game() graph.GameResolver { return &gameResolver{r} }

type gameResolver struct{ *Resolver }

// IsMatch is the resolver for the isMatch field.
func (r *gameResolver) IsMatch(ctx context.Context, obj *model.Game) (bool, error) {
	return entities.GameIsMatch(ctx, r.CustomerAPI, obj)
}
