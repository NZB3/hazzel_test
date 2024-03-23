package projectscontroller

import (
	"context"
	"crudserver/internal/app/models"
	"crudserver/internal/logger"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

type projectsStorage interface {
	SaveProject(ctx context.Context, project *models.Project) error
	GetProjects(ctx context.Context) ([]*models.Project, error)
	GetProjectByID(ctx context.Context, id int) (*models.Project, error)
	UpdateProject(ctx context.Context, project *models.Project) error
	DeleteProject(ctx context.Context, id int) error
}

type controller struct {
	log     logger.Logger
	storage projectsStorage
}

type projectsRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type projectsResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func New(log logger.Logger, storage projectsStorage) *controller {
	return &controller{
		log:     log,
		storage: storage,
	}
}

func (c *controller) AddRoutsToApp(app *fiber.App) {
	app.Get("/projects", c.getAllHandler())
	app.Get("/project", c.getOneHandler())
	app.Post("/project", c.createHandler())
	app.Patch("/project", c.updateHandler())
	app.Delete("/project", c.deleteHandler())
}

func (c *controller) notFoundError(details ...any) string {
	errRes, err := json.Marshal(&models.ErrorResponse{
		Code:    3,
		Message: "errors.project.notFound",
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
