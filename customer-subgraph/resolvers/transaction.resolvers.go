package resolvers

import (
	"context"
	"customer-subgraph/graph"
	"customer-subgraph/graph/model"
	"customer-subgraph/resolvers/entities"
)

// Transaction returns graph.TransactionResolver implementation.
func (r *Resolver) Transaction() graph.TransactionResolver { return &transactionResolver{r} }

type transactionResolver struct{ *Resolver }

// Bets is the resolver for the bets field.
func (r *transactionResolver) Bets(ctx context.Context, obj *model.Transaction) ([]*model.Bet, error) {
	return entities.TransactionBets(ctx, r.CustomerAPI, obj)
}
