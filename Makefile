include .env
export

MIGRATIONS_DIR=migrations

.PHONY: migration-new migration-new-go migration-up migration-down

migration-new:
	@read -p "Enter migration (sql) name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql

migration-new-go:
	@read -p "Enter migration (go) name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name go

migration-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

migration-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

migration-status:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) status

