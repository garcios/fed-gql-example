package entities

import (
	"context"
	"customer-subgraph/datasources"
	"customer-subgraph/graph/model"
)

// FindGameByID resolves a Game entity by its unique ID.
func FindGameByID(ctx context.Context, api *datasources.MockAPI, id string) (*model.Game, error) {
	return &model.Game{
		ID: id,
	}, nil
}

// GameIsMatch resolves the isMatch field for a Game entity.
func GameIsMatch(ctx context.Context, api *datasources.MockAPI, obj *model.Game) (bool, error) {
	switch obj.MarketTypeID {
	case "mt_match_winner", "mt_point_spread", "mt_total_points":
		return true, nil
	case "multi_runner":
		return false, nil
	default:
		return false, nil
	}

}
