package app

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	app := fiber.New()

	Config(app)
	InitDependencies(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	app.Listen(fmt.Sprintf(":%s", port))
}