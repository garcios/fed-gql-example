package entities

import (
	"context"
	"customer-subgraph/graph/model"
)

// FindCustomerByID resolves a Customer entity by its unique ID.
func FindCustomerByID(ctx context.Context, id string) (*model.Customer, error) {
	return &model.Customer{
		ID: id,
	}, nil
}

// CustomerTransactions resolves the transactions field for a Customer entity.
func CustomerTransactions(ctx context.Context, obj *model.Customer) ([]*model.Transaction, error) {
	return []*model.Transaction{
		{
			CustomerID: obj.ID,
			Bets: []*model.Bet{
				{
					ID:     "bet_001",
					Amount: 120.50,
					Game:   &model.Game{ID: "game_777"},
				},
				{
					ID:     "bet_002",
					Amount: 45.00,
					Game:   &model.Game{ID: "game_888"},
				},
			},
		},
	}, nil
}
