package resolvers

import "game-subgraph/datasources"

type Resolver struct {
	GameAPI *datasources.MockAPI
}
