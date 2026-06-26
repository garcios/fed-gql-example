package entities

import (
	"context"
	"game-subgraph/datasources"
	"game-subgraph/graph/model"
)

// FindGameByID resolves a Game entity by its unique ID.
func FindGameByID(ctx context.Context, api *datasources.MockAPI, id string) (*model.Game, error) {
	return api.GetGame(ctx, id)
}
