DB_URL=postgres://binhbeng:1998@localhost:5432/goapp?sslmode=disable
MIGRATIONS_DIR=./migrations

# --- TARGETS ---
.PHONY: migrate-up migrate-down migrate-new migrate-force

# Create a new migration
migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

# Apply all pending migrations
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Rollback 1 migration
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Force version (trong trường hợp bị lỗi dirty)
migrate-force:
	@read -p "Enter version: " v; \
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$v