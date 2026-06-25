package graph

import (
	"context"
	"game-subgraph/graph/model"
)

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	return &model.Game{
		ID:           id,
		Name:         "Premier League Football Match",
		MarketTypeID: "mt_head_to_head",
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
