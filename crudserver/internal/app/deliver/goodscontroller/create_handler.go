package goodscontroller

import (
	"crudserver/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func (c *controller) createHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c.log.Info("Create good")
		var goodsReq goodsRequest
		err := ctx.BodyParser(&goodsReq)
		if err != nil {
			c.log.Errorf("Failed to parse request: %s", err)
			return fiber.NewError(fiber.StatusBadRequest, c.badRequestError())
		}

		good := &models.Good{
			ProjectID:   goodsReq.ProjectID,
			Name:        goodsReq.Name,
			Description: goodsReq.Description,
		}

		err = c.storage.SaveGood(ctx.Context(), good)
		if err != nil {
			c.log.Errorf("Failed to save good: %s", err)
			return fiber.NewError(fiber.StatusInternalServerError, c.internalServerError())
		}

		return ctx.SendStatus(fiber.StatusCreated)
	}
}
