.PHONY: start-customer start-game start-all help

help:
	@echo "Available commands:"
	@echo "  make start-customer - Start the Customer Subgraph (Port 8081)"
	@echo "  make start-game     - Start the Game Subgraph (Port 8082)"
	@echo "  make start-all      - Start both Customer and Game Subgraphs concurrently"

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
