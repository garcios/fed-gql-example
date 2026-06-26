package queries

import (
	"context"
	"game-subgraph/datasources"
	"game-subgraph/graph/model"
)

// Game resolves the query for a single Game by ID.
func Game(ctx context.Context, api *datasources.MockAPI, id string) (*model.Game, error) {
	return api.GetGame(ctx, id)
}
