.PHONY: start-customer start-game start-all help generate generate-customer generate-game rover-dev

help:
	@echo "Available commands:"
	@echo "  make start-customer - Start the Customer Subgraph (Port 8081)"
	@echo "  make start-game     - Start the Game Subgraph (Port 8082)"
	@echo "  make start-all      - Start both Customer and Game Subgraphs concurrently"
	@echo "  make rover-dev      - Start the Rover development server"
	@echo "  make generate       - Generate gqlgen code for all subgraphs"
	@echo "  make generate-customer - Generate gqlgen code for Customer Subgraph"
	@echo "  make generate-game     - Generate gqlgen code for Game Subgraph"

start-customer:
	@echo "Starting Customer Subgraph..."
	cd customer-subgraph && go run server.go

start-game:
	@echo "Starting Game Subgraph..."
	cd game-subgraph && go run server.go

start-all:
	@echo "Starting all subgraphs concurrently..."
	$(MAKE) -j 2 start-customer start-game

rover-dev:
	rover dev --supergraph-config ./supergraph-config.yaml

generate: generate-customer generate-game

generate-customer:
	@echo "Generating gqlgen code for Customer Subgraph..."
	cd customer-subgraph && go run github.com/99designs/gqlgen generate

generate-game:
	@echo "Generating gqlgen code for Game Subgraph..."
	cd game-subgraph && go run github.com/99designs/gqlgen generate
