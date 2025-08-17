package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zulfahmidev/shorturl/pkg/constant"
)

func RunServer() {
	app := fiber.New()

	SetupResource(app)
	InitDependencies(app)

	app.Listen(fmt.Sprintf(":%d", constant.APP_PORT))
}