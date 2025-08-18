package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	app := fiber.New()

	Config(app)
	InitDependencies(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}