package queries

import (
	"context"
	"game-subgraph/graph/model"
)

// Game resolves the query for a single Game by ID.
func Game(ctx context.Context, id string) (*model.Game, error) {
	return &model.Game{
		ID:           id,
		Name:         "Premier League Football Match",
		MarketTypeID: "mt_head_to_head",
	}, nil
}
