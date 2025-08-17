package handler

import "github.com/gofiber/fiber/v2"

type IURLService interface {
	
}

type URLHandler struct {
	urlService *IURLService
}

func NewUrlHandler(app *fiber.App, urlService IURLService) *URLHandler {
	handler := &URLHandler{
		urlService: &urlService,
	}

	app.Post("/shorten", handler.CreateShortenURL)
	app.Get("/:shortToken", handler.GetRedirectURL)
	app.Get("/stats/:shortToken", handler.GetStatsURL)

	return handler
}

func (h *URLHandler) CreateShortenURL(c *fiber.Ctx) error {

	return nil
}

func (h *URLHandler) GetRedirectURL(c *fiber.Ctx) error {
	
	return nil
}

func (h *URLHandler) GetStatsURL(c *fiber.Ctx) error {
	
	return nil
}
