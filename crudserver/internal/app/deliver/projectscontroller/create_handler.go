package projectscontroller

import (
	"crudserver/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func (c *controller) createHandler() fiber.Handler {
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

		err = c.storage.SaveProject(ctx.Context(), project)
		if err != nil {
			c.log.Errorf("Failed to save project: %s", err)
			return fiber.NewError(fiber.StatusInternalServerError, c.internalServerError())
		}

		return ctx.SendStatus(fiber.StatusCreated)
	}
}
