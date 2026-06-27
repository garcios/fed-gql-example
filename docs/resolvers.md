# GraphQL Resolvers: Struct Embedding & Sub-Resolvers

In the Federated GraphQL implementation, we use **gqlgen**'s split resolver pattern. This document explains the structure and usage of sub-resolvers, specifically focusing on the pattern:

```go
type betResolver struct{ *Resolver }
```

---

## 1. Struct Embedding in Go

By declaring `*Resolver` as an anonymous field inside `betResolver`, Go performs **struct embedding** (specifically, embedding a pointer to the root `Resolver` struct).

* **Field & Method Promotion**: All fields and methods of the base [Resolver](file:///Users/oscargarcia/workspace/fed-gql-example/customer-subgraph/resolvers/resolver.go#L5-L7) are promoted directly to [betResolver](file:///Users/oscargarcia/workspace/fed-gql-example/customer-subgraph/resolvers/bet.resolvers.go#L13).
* **Direct Access**: Methods on `betResolver` can access dependencies (like `CustomerAPI`) directly via `r.CustomerAPI` instead of having to go through a nested field like `r.Resolver.CustomerAPI`.

---

## 2. Implementing gqlgen's `BetResolver` Interface

When a GraphQL type defines fields that require custom resolver logic (for example, fetching a nested `game` object for a `Bet`), **gqlgen** generates a dedicated resolver interface (e.g., `graph.BetResolver`).

To satisfy this interface:
1. **The Factory Method**: The root `Resolver` implements a method to return the sub-resolver:
   ```go
   func (r *Resolver) Bet() graph.BetResolver { return &betResolver{r} }
   ```
   This constructs a new `betResolver` and injects the current root resolver context `r`.
2. **Field Resolvers**: The `betResolver` struct then implements the specific field-level resolver methods defined in the interface (e.g., `Game` and `Amount` on [bet.resolvers.go](file:///Users/oscargarcia/workspace/fed-gql-example/customer-subgraph/resolvers/bet.resolvers.go#L16-L23)).

---

## 3. Benefits of This Design

* **Access to Shared Context**: It shares database connections, mock APIs, and state (from the main `Resolver`) down to all sub-resolvers.
* **Separation of Concerns**: It prevents the root `Resolver` namespace from being cluttered with every field resolver method, keeping each type's resolver methods cleanly grouped in their respective files.
