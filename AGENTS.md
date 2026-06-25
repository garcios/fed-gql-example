# AGENTS.md: Federated GraphQL Implementation Blueprint (Golang + gqlgen)

This document contains automated orchestration instructions for creating a fully functional Federated GraphQL architecture using **Golang**, **gqlgen (v2 federation)**, and a decoupled **Datasources structure**.

---

## 1. System & Technology Stack
- **Language**: Go 1.23+
- **GraphQL Engine**: `github.com/99designs/gqlgen` (Configured for Apollo Federation v2)
- **Monorepo Directory Layout**:
  ```text
  / (Root)
  ├── customer-subgraph/
  │   ├── datasources/         # Mock database / API layer for Customers, Transactions, Bets
  │   ├── graph/               # GraphQL delivery layer (schemas, generated code, resolvers)
  │   ├── server.go            # Entrypoint
  │   └── gqlgen.yml
  ├── game-subgraph/
  │   ├── datasources/         # Mock database / API layer for Games and MarketTypes
  │   ├── graph/               # GraphQL delivery layer (schemas, generated code, resolvers)
  │   ├── server.go            # Entrypoint
  │   └── gqlgen.yml
  └── AGENTS.md
  ```

## 2. Federated Data Architecture

```text
┌────────────────────────┐
│  Apollo Router/Gateway  │
└───────────┬────────────┘
            │
┌───────────┴─────────────────────┐
▼                                 ▼
┌───────────────────────┐         ┌───────────────────────┐
│   Customer Subgraph   │         │     Game Subgraph     │
│      (Port 8081)      │         │      (Port 8082)      │
├───────────────────────┤         ├───────────────────────┤
│ type Customer @key    │         │ type Game @key        │
│ type Transaction      │         │ type MarketType       │
│ type Bet              │         │                       │
│ type Game @external   │         │                       │
└───────────────────────┘         └───────────────────────┘
```

### Data Boundaries

- **Customer Subgraph**: Owns personal customer attributes, transactions, and individual bets. It holds a reference pointer to `Game` via its unique identifier (`id`), treating `Game` as an external entity stub.
- **Game Subgraph**: Owns the single source of truth for games and their accompanying metadata (e.g., name and `marketTypeID`). It exposes entity resolvers (`FindGameByID`) so the gateway can hydrate game data requested from within a customer context.

---

## 3. Automation Implementation Guide

Follow these sequential blocks to initialize, configure, and code the full system.

### Step 3.1: Workspace Initialization

Run the following shell setup to build out the directory modules and initialize the basic code structures.

```bash
# Initialize Customer Subgraph
mkdir -p customer-subgraph/graph
cd customer-subgraph
go mod init customer-subgraph
cat << 'EOF' > tools.go
//go:build tools
package tools
import _ "github.com/99designs/gqlgen"
EOF
go get github.com/99designs/gqlgen
cd ..

# Initialize Game Subgraph
mkdir -p game-subgraph/graph
cd game-subgraph
go mod init game-subgraph
cat << 'EOF' > tools.go
//go:build tools
package tools
import _ "github.com/99designs/gqlgen"
EOF
go get github.com/99designs/gqlgen
cd ..
```

### Step 3.2: Configuration Specifications

#### Customer `gqlgen.yml`

```yaml
# customer-subgraph/gqlgen.yml
schema:
  - graph/*.graphqls
exec:
  filename: graph/generated.go
  package: graph
model:
  filename: graph/model/models_gen.go
  package: model
resolver:
  layout: follow-schema
  dir: graph
  package: graph
federation:
  version: 2
```

#### Game `gqlgen.yml`

```yaml
# game-subgraph/gqlgen.yml
schema:
  - graph/*.graphqls
exec:
  filename: graph/generated.go
  package: graph
model:
  filename: graph/model/models_gen.go
  package: model
resolver:
  layout: follow-schema
  dir: graph
  package: graph
federation:
  version: 2
```

### Step 3.3: GraphQL Schema Layouts

#### Customer Schema (`customer-subgraph/graph/schema.graphqls`)

```graphql
extend schema
  @link(url: "https://specs.apollo.dev/federation/v2.0",
        import: ["@key", "@external"])

type Customer @key(fields: "id") {
  id: ID!
  transactions: [Transaction!]!
}

type Transaction {
  customerID: ID!
  bets: [Bet!]!
}

type Bet {
  id: ID!
  game: Game!
  amount: Float!
}

type Game @key(fields: "id") {
  id: ID! @external
}

type Query {
  customer(id: ID!): Customer
}
```

#### Game Schema (`game-subgraph/graph/schema.graphqls`)

```graphql
extend schema
  @link(url: "https://specs.apollo.dev/federation/v2.0",
        import: ["@key"])

type Game @key(fields: "id") {
  id: ID!
  name: String!
  marketTypeID: ID!
}

type Query {
  game(id: ID!): Game
}
```

### Step 3.4: Generation Code Triggers

Execute these statements to generate boilerplate code engines within both contexts.

```bash
cd customer-subgraph
go run github.com/99designs/gqlgen generate
cd ../game-subgraph
go run github.com/99designs/gqlgen generate
cd ..
```

### Step 3.5: Resolver & Entrypoint Implementations

#### Customer Query Resolver Implementation (`customer-subgraph/graph/schema.resolvers.go`)

```go
package graph

import (
	"context"
	"customer-subgraph/graph/model"
)

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	return &model.Customer{
		ID: id,
		Transactions: []*model.Transaction{
			{
				CustomerID: id,
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
		},
	}, nil
}
```

#### Customer Server Entrypoint (`customer-subgraph/server.go`)

```go
package main

import (
	"customer-subgraph/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for Customer Subgraph playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

#### Game Standard & Entity Resolver Implementation

##### Schema Resolvers (`game-subgraph/graph/schema.resolvers.go`)

```go
package graph

import (
	"context"
	"game-subgraph/graph/model"
)

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	return &model.Game{
		ID:           id,
		Name:         "Premier League Football Match",
		MarketTypeID: "mt_head_to_head",
	}, nil
}
```

##### Entity Resolvers (`game-subgraph/graph/entity.resolvers.go`)

```go
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
```

#### Game Server Entrypoint (`game-subgraph/server.go`)

```go
package main

import (
	"game-subgraph/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for Game Subgraph playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

---

## 4. Execution & Verification Flow

To run the complete federated architecture, spin up both subgraphs and point your Apollo Router or Gateway config to them:

**Start Customer Subgraph:**

```bash
cd customer-subgraph && go run server.go
```

**Start Game Subgraph:**

```bash
cd game-subgraph && go run server.go
```

### Composite Gateway Query Verification

When requesting across the supergraph gateway layer, run the following unified execution block to test nested federation hydration:

```graphql
query GetCustomerActivityRecord {
  customer(id: "cust_usr_901") {
    transactions {
      customerID
      bets {
        id
        amount
        game {
          id
          name
          marketTypeID
        }
      }
    }
  }
}
```