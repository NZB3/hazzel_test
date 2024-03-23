package goodscontroller

import (
	"context"
	"crudserver/internal/app/models"
	"crudserver/internal/logger"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

type goodsStorage interface {
	SaveGood(ctx context.Context, good *models.Good) error
	GetGoods(ctx context.Context) ([]*models.Good, error)
	GetGoodByID(ctx context.Context, id int) (*models.Good, error)
	UpdateGood(ctx context.Context, good *models.Good) error
	DeleteGood(ctx context.Context, id int) error
}

type goodsRequest struct {
	ProjectID   int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type goodsResponse struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
}

type controller struct {
	log     logger.Logger
	storage goodsStorage
}

func New(log logger.Logger, storage goodsStorage) *controller {
	return &controller{
		log:     log,
		storage: storage,
	}
}

func (c *controller) AddRoutsToApp(app *fiber.App) {
	app.Get("/goods", c.getAllHandler())
	app.Get("/good", c.getOneHandler())
	app.Post("/good", c.createHandler())
	app.Patch("/good", c.updateHandler())
	app.Delete("/good", c.deleteHandler())
}

func (c *controller) notFoundError(details ...any) string {
	errRes, err := json.Marshal(&models.ErrorResponse{
		Code:    3,
		Message: "errors.good.notFound",
		Details: details,
	})

	if err != nil {
		c.log.Error("Failed to create error response")
	}

	return string(errRes)
}

func (c *controller) badRequestError(details ...any) string {
	errRes, err := json.Marshal(&models.ErrorResponse{
		Code:    2,
		Message: "errors.badRequest",
		Details: details,
	})

	if err != nil {
		c.log.Error("Failed to create error response")
	}

	return string(errRes)
}

func (c *controller) internalServerError(details ...any) string {
	errRes, err := json.Marshal(&models.ErrorResponse{
		Code:    1,
		Message: "errors.internal",
		Details: details,
	})

	if err != nil {
		c.log.Error("Failed to create error response")
	}

	return string(errRes)
}
