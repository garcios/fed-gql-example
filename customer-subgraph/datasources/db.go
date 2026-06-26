package datasources

import (
	"context"
	"customer-subgraph/graph/model"
)

// MockAPI simulates a datasource/database for the Customer Subgraph.
type MockAPI struct {
	customers    map[string]*model.Customer
	transactions map[string][]*model.Transaction
	bets         map[string][]*model.Bet
}

// NewMockAPI creates and populates a new MockAPI instance with mock data.
func NewMockAPI() *MockAPI {
	return &MockAPI{
		customers: map[string]*model.Customer{
			"cust_usr_901": {ID: "cust_usr_901"},
			"cust_usr_902": {ID: "cust_usr_902"},
		},
		transactions: map[string][]*model.Transaction{
			"cust_usr_901": {
				{CustomerID: "cust_usr_901"},
			},
			"cust_usr_902": {
				{CustomerID: "cust_usr_902"},
			},
		},
		bets: map[string][]*model.Bet{
			"cust_usr_901": {
				{ID: "bet_001", Amount: 120.50, GameID: "game_777"},
				{ID: "bet_002", Amount: 45.00, GameID: "game_888"},
				{ID: "bet_003", Amount: 100.00, GameID: "game_999"},
			},
			"cust_usr_902": {
				{ID: "bet_004", Amount: 250.00, GameID: "game_777"},
			},
		},
	}
}

// GetCustomer retrieves a Customer by ID.
func (api *MockAPI) GetCustomer(ctx context.Context, id string) (*model.Customer, error) {
	if c, exists := api.customers[id]; exists {
		return c, nil
	}
	// GraphQL typically expects nil if an entity is not found rather than an error,
	// but we can also return a default initialized customer or nil.
	return nil, nil
}

// GetTransactions retrieves Transactions for a given Customer ID.
func (api *MockAPI) GetTransactions(ctx context.Context, customerID string) ([]*model.Transaction, error) {
	if t, exists := api.transactions[customerID]; exists {
		return t, nil
	}
	return []*model.Transaction{}, nil
}

// GetBets retrieves Bets for a given Customer ID.
func (api *MockAPI) GetBets(ctx context.Context, customerID string) ([]*model.Bet, error) {
	if b, exists := api.bets[customerID]; exists {
		return b, nil
	}
	return []*model.Bet{}, nil
}
