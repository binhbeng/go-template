# DB_URL=postgres://binhbeng:1998@localhost:5432/goapp?sslmode=disable
# MIGRATIONS_DIR=migrations
MIGRATION_DIRS = internal/db/migrations

# .PHONY: migration-new migration-new-go migration-up migration-down

# migration-new:
# 	@read -p "Enter migration (sql) name: " name; \
# 	goose -dir $(MIGRATIONS_DIR) create $$name sql

# migration-new-go:
# 	@read -p "Enter migration (go) name: " name; \
# 	goose -dir $(MIGRATIONS_DIR) create $$name go

# migration-up:
# 	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

# migration-down:
# 	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

# migration-status:
# 	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) status

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIRS) -seq $(NAME)
