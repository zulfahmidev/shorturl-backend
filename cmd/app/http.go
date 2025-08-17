package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zulfahmidev/shorturl/cmd/config"
)

func RunServer() {
	app := fiber.New()

	SetupResource(app)
	InitDependencies(app)

	app.Listen(fmt.Sprintf(":%d", config.APP_PORT))
}