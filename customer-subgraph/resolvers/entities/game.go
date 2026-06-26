package entities

import (
	"context"
	"customer-subgraph/graph/model"
)

// FindGameByID resolves a Game entity by its unique ID.
func FindGameByID(ctx context.Context, id string) (*model.Game, error) {
	return &model.Game{
		ID: id,
	}, nil
}

// GameIsBettable resolves the isBettable field for a Game entity.
func GameIsBettable(ctx context.Context, obj *model.Game) (bool, error) {
	// A sporting game is considered bettable if the market type is not standard or empty.
	return obj.MarketTypeID != "" && obj.MarketTypeID != "mt_standard", nil
}
