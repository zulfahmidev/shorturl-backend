package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zulfahmidev/shorturl/internal/dto"
	"github.com/zulfahmidev/shorturl/pkg/logger"
	"github.com/zulfahmidev/shorturl/pkg/rest"
	"github.com/zulfahmidev/shorturl/pkg/validator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IURLService interface {
	GenerateShortURL(url string, idempotent bool) (*dto.URLData, error)
	GetStatsURL(token string) (*dto.URLData, error)
}

type URLHandler struct {
	urlService IURLService
}

func NewUrlHandler(app *fiber.App, urlService IURLService) *URLHandler {
	handler := &URLHandler{
		urlService: urlService,
	}

	app.Post("/shorten", handler.GenerateShortURL)
	app.Get("/:shortToken", handler.GetRedirectURL)
	app.Get("/stats/:shortToken", handler.GetStatsURL)

	return handler
}

func (h *URLHandler) GenerateShortURL(c *fiber.Ctx) error {
	type GenerateShortURLRequest struct {
		URL        *string `form:"url"`
		Idempotent *bool   `form:"idempotent"`
	}

	var req GenerateShortURLRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(rest.BaseResponse{
			Message: "something went wrong",
			Body:    nil,
		})
	}

	if req.URL == nil {
		err := errors.New("missing url")
		return c.Status(http.StatusBadRequest).JSON(rest.BaseResponse{
			Message: err.Error(),
			Body:    nil,
		})
	}

	if !validator.IsValidURL(*req.URL) {
		err := errors.New("invalid url format")
		return c.Status(http.StatusUnprocessableEntity).JSON(rest.BaseResponse{
			Message: err.Error(),
			Body:    nil,
		})
	}

	idempotent := true
	if req.Idempotent != nil {
		idempotent = *req.Idempotent
	}

	data, err := h.urlService.GenerateShortURL(*req.URL, idempotent)
	if err != nil {
		logger.Default().Error(err.Error(), zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(rest.BaseResponse{
			Message: "something went wrong",
			Body:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(rest.BaseResponse{
		Message: "short url generated successfully",
		Body:    data,
	})
}

func (h *URLHandler) GetRedirectURL(c *fiber.Ctx) error {

	return nil
}

func (h *URLHandler) GetStatsURL(c *fiber.Ctx) error {
	shortToken := c.Params("shortToken")
	data, err := h.urlService.GetStatsURL(shortToken)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(rest.BaseResponse{
				Message: "stats request for invalid token",
				Body:    nil,
			})
		}

		logger.Default().Error(err.Error(), zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(rest.BaseResponse{
			Message: "something went wrong",
			Body:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(rest.BaseResponse{
		Message: "statistics loaded successfully",
		Body:    data,
	})
}
