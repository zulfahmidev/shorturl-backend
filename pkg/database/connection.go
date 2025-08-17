package database

import (
	"fmt"
	"log"

	"github.com/zulfahmidev/shorturl/pkg/constant"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() {
	var (
		dsn string
		conn *gorm.DB
		err error
	)

	switch constant.DATABASE_DRIVER {
	case "postgres":
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			constant.DATABASE_USER,
			constant.DATABASE_PASSWORD,
			constant.DATABASE_HOST,
			constant.DATABASE_PORT,
			constant.DATABASE_NAME,
		)
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect postgres: %v", err)
		}

		conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
		conn.Exec(`CREATE EXTENSION IF NOT EXISTS postgis`)

	case "sqlite3":
		dsn = "./shorturl.db"
		conn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect sqlite: %v", err)
		}

	default:
		log.Fatalf("unsupported driver: %s", constant.DATABASE_DRIVER)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
