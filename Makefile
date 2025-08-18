DB_URL = "postgres://trovary:password@localhost:5432/db_shorturl?sslmode=disable"
MIGRATIONS = db/migration
MIGRATE = migrate -path $(MIGRATIONS) -database $(DB_URL)

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
