package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
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

	dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}

	conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	conn.Exec(`CREATE EXTENSION IF NOT EXISTS postgis`)

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
