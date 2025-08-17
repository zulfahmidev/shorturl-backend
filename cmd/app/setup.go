package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulfahmidev/shorturl/pkg/cors"
	"github.com/zulfahmidev/shorturl/pkg/database"
	"github.com/zulfahmidev/shorturl/pkg/logger"
	"github.com/zulfahmidev/shorturl/pkg/validator"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func SetupResource(app *fiber.App) {
	validator.InitValidator()
	logger.InitLogger()
	cors.InitCors(app)
	database.InitDB()
}