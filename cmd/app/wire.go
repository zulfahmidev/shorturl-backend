package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulfahmidev/shorturl/internal/handler"
	"github.com/zulfahmidev/shorturl/internal/repository"
	"github.com/zulfahmidev/shorturl/internal/service"
)

func InitDependencies(app *fiber.App) {
	urlRepository := repository.NewURLRepository()

	urlService := service.NewURLService(urlRepository)
	
	handler.NewUrlHandler(app, urlService)
}