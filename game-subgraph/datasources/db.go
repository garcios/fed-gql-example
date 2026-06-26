package datasources

import (
	"context"
	"game-subgraph/graph/model"
)

// MockAPI simulates a datasource/database for the Game Subgraph.
type MockAPI struct {
	games map[string]*model.Game
}

// NewMockAPI creates and populates a new MockAPI instance with mock data.
func NewMockAPI() *MockAPI {
	return &MockAPI{
		games: map[string]*model.Game{
			"game_777": {
				ID:           "game_777",
				Name:         "Championship Basketball Finals",
				MarketTypeID: "mt_point_spread",
			},
			"game_888": {
				ID:           "game_888",
				Name:         "World Cup Tennis Open",
				MarketTypeID: "mt_match_winner",
			},
			"game_999": {
				ID:           "game_999",
				Name:         "F1 racing",
				MarketTypeID: "multi_runner",
			},
			"game_666": {
				ID:           "game_666",
				Name:         "Premier League Football Match",
				MarketTypeID: "mt_head_to_head",
			},
		},
	}
}

// GetGame retrieves a Game by ID, falling back to a default sport event if not found.
func (api *MockAPI) GetGame(ctx context.Context, id string) (*model.Game, error) {
	if g, exists := api.games[id]; exists {
		return g, nil
	}
	return &model.Game{
		ID:           id,
		Name:         "Generic Sporting Event",
		MarketTypeID: "mt_standard",
	}, nil
}
