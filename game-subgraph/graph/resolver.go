package graph

import "game-subgraph/datasources"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	GameAPI *datasources.MockAPI
}
