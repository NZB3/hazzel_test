package projectscontroller

import (
	"crudserver/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"strconv"
)

func (c *controller) getOneHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idQuery := ctx.Params("id")
		if idQuery == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Bad request")
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Bad request")
		}

		project, err := c.storage.GetProjectByID(ctx.Context(), id)
		if err != nil {
			c.log.Errorf("Failed to get project: %s", err)
			return fiber.NewError(fiber.StatusInternalServerError, c.internalServerError())
		}

		return ctx.JSON(&projectsResponse{
			ID:        project.ID,
			Name:      project.Name,
			CreatedAt: project.CreatedAt,
		})
	}
}

func (c *controller) getAllHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		projects, err := c.storage.GetProjects(ctx.Context())
		if err != nil {
			c.log.Errorf("Failed to get projects: %s", err)
			return fiber.NewError(fiber.StatusInternalServerError, c.internalServerError())
		}

		return ctx.JSON(lo.Map(projects, func(project *models.Project, _ int) projectsResponse {
			return projectsResponse{
				ID:        project.ID,
				Name:      project.Name,
				CreatedAt: project.CreatedAt,
			}
		}))
	}
}
