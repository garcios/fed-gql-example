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
		},
	}, nil
}

// TransactionBets resolves the bets field for a Transaction entity.
func TransactionBets(ctx context.Context, obj *model.Transaction) ([]*model.Bet, error) {
	return []*model.Bet{
		{
			ID:     "bet_001",
			Amount: 120.50,
			GameID: "game_777",
		},
		{
			ID:     "bet_002",
			Amount: 45.00,
			GameID: "game_888",
		},
	}, nil
}

// BetGame resolves the game field for a Bet entity.
func BetGame(ctx context.Context, obj *model.Bet) (*model.Game, error) {
	return &model.Game{
		ID: obj.GameID,
	}, nil
}
