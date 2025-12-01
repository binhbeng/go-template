DB_URL=postgres://binhbeng:1998@localhost:5432/goapp?sslmode=disable
MIGRATIONS_DIR=migrations

# --- TARGETS ---
# .PHONY: migrate-up migrate-down migrate-new migrate-force migrate-check

# # Create a new migration
# migrate-new:
# 	@read -p "Enter migration name: " name; \
# 	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

# # Apply all pending migrations
# migrate-up:
# 	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# # Rollback 1 migration
# migrate-down:
# 	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# # Force version (in case error: Dirty database)
# migrate-force:
# 	@read -p "Enter version: " v; \
# 	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$v

# migrate-check:
# 	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

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

#####
proto:
	protoc --go_out=. --go-grpc_out=. internal/proto/*.proto