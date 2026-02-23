DB_PATH ?= ./data/pedro.db
MIGRATIONS_PATH ?= ./migrations
DB_URL := sqlite3://$(DB_PATH)

.PHONY: migrate-up migrate-down migrate-version migrate-create

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

# Usage: make migrate-create name=add_last_cleaned_at
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)