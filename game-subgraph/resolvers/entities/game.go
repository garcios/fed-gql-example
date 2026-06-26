package entities

import (
	"context"
	"game-subgraph/graph/model"
)

// FindGameByID resolves a Game entity by its unique ID.
func FindGameByID(ctx context.Context, id string) (*model.Game, error) {
	var name string
	var marketTypeID string

	switch id {
	case "game_777":
		name = "Championship Basketball Finals"
		marketTypeID = "mt_point_spread"
	case "game_888":
		name = "World Cup Tennis Open"
		marketTypeID = "mt_match_winner"
	case "game_999":
		name = "F1 racing"
		marketTypeID = "multi_runner"
	default:
		name = "General Sporting Event"
		marketTypeID = "mt_standard"
	}

	return &model.Game{
		ID:           id,
		Name:         name,
		MarketTypeID: marketTypeID,
	}, nil
}
