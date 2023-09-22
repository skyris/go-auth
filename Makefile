-include .env
DB_STRING="postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)?sslmode=$(DB_SSLMODE)"

up:
	@docker compose up --build
.PHONY: up

down:
	@docker compose down
.PHONY: down

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_STRING) up
.PHONY: migrate-up
