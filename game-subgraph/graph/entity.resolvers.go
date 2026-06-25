package graph

import (
	"context"
	"game-subgraph/graph/model"
)

func (r *entityResolver) FindGameByID(ctx context.Context, id string) (*model.Game, error) {
	var name string
	var marketTypeID string

	switch id {
	case "game_777":
		name = "Championship Basketball Finals"
		marketTypeID = "mt_point_spread"
	case "game_888":
		name = "World Cup Tennis Open"
		marketTypeID = "mt_match_winner"
	default:
		name = "Generic Sporting Event"
		marketTypeID = "mt_standard"
	}

	return &model.Game{
		ID:           id,
		Name:         name,
		MarketTypeID: marketTypeID,
	}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
