db ?= postgres

POSTGRES_DSN = "postgres://user:password@localhost:5432/dbname?sslmode=disable"

SQLITE_FILE = "./data.db"
SQLITE_DSN = "sqlite3://$(SQLITE_FILE)"

ifeq ($(db),postgres)
	DB_URL = $(POSTGRES_DSN)
else ifeq ($(db),sqlite)
	DB_URL = $(SQLITE_DSN)
else
	$(error DB harus postgres atau sqlite)
endif

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
