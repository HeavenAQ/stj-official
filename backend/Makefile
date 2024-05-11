# Check if .env file exists
ifeq ($(wildcard app.env),)
    $(info No .env file found, proceeding without it...)
else
    include app.env
endif

migrate_up: 
	@echo "Migrating up..."
	migrate -path $(DB_MIGRATION_DIR) -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" -verbose up
	@echo "Migration up completed"

migrate_down:
	@echo "Migrating down..."
	migrate -path $(DB_MIGRATION_DIR) -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" -verbose down
	@echo "Migration down completed"

migrate_create:
	@echo "Creating migration..."
	migrate create -ext sql -dir $(DB_MIGRATION_DIR) -seq $(DB_MIGRATION_NAME)
	@echo "Migration created"

sqlc:
	@echo "Generating sqlc..."
	sqlc generate
	@echo "sqlc generated"

test: 
	go test -v -cover ./... 
