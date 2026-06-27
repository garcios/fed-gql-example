package resolvers

import "customer-subgraph/datasources"

type Resolver struct {
	CustomerAPI *datasources.MockAPI
}
