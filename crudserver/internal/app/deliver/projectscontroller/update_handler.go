package projectscontroller

import (
	"crudserver/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func (c *controller) updateHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var projectsReq projectsRequest
		err := ctx.BodyParser(&projectsReq)
		if err != nil {
			c.log.Errorf("Failed to parse request: %s", err)
			return fiber.NewError(fiber.StatusBadRequest, c.badRequestError())
		}

		project := &models.Project{
			ID:   projectsReq.ID,
			Name: projectsReq.Name,
		}

		err = c.storage.UpdateProject(ctx.Context(), project)
		if err != nil {
			c.log.Errorf("Failed to update project: %s", err)
			return fiber.NewError(fiber.StatusNotFound, c.notFoundError())
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
