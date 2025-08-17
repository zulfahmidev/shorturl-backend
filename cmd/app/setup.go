package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zulfahmidev/shorturl/cmd/config"

    _ "github.com/lib/pq" 
    _ "github.com/mattn/go-sqlite3"
)

func SetupResource(app *fiber.App) {
	var dsn string

    switch config.DATABASE_DRIVER {
		case "postgres":
			dsn = fmt.Sprintf(
				"postgres://%s:%s@%s:%d/%s?sslmode=disable",
				config.DATABASE_USER,
				config.DATABASE_PASSWORD,
				config.DATABASE_HOST,
				config.DATABASE_PORT,
				config.DATABASE_NAME,
			)
		case "sqlite3":
			dsn = "./shorturl.db"
		default:
			log.Fatalf("unsupported driver: %s", config.DATABASE_DRIVER)
    }

    db, err := sql.Open(config.DATABASE_DRIVER, dsn)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()
}