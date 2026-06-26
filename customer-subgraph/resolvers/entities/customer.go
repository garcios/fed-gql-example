package entities

import (
	"context"
	"customer-subgraph/datasources"
	"customer-subgraph/graph/model"
)

// FindCustomerByID resolves a Customer entity by its unique ID.
func FindCustomerByID(ctx context.Context, api *datasources.MockAPI, id string) (*model.Customer, error) {
	return api.GetCustomer(ctx, id)
}

// CustomerTransactions resolves the transactions field for a Customer entity.
func CustomerTransactions(ctx context.Context, api *datasources.MockAPI, obj *model.Customer) ([]*model.Transaction, error) {
	return api.GetTransactions(ctx, obj.ID)
}

// TransactionBets resolves the bets field for a Transaction entity.
func TransactionBets(ctx context.Context, api *datasources.MockAPI, obj *model.Transaction) ([]*model.Bet, error) {
	return api.GetBets(ctx, obj.CustomerID)
}

// BetGame resolves the game field for a Bet entity.
func BetGame(ctx context.Context, api *datasources.MockAPI, obj *model.Bet) (*model.Game, error) {
	return &model.Game{
		ID: obj.GameID,
	}, nil
}
