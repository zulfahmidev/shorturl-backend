include .env
export $(shell sed 's/=.*//' .env)

DB_URL = "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable"
MIGRATIONS = db/migration
MIGRATE = migrate -path $(MIGRATIONS) -database $(DB_URL)

print-env:
	@echo "User: $(DB_USER), Pass: $(DB_PASS)"

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-reset:
	$(MIGRATE) down -all
	$(MIGRATE) up

migrate-new:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-new name=nama_migration"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS) -seq $(name)
