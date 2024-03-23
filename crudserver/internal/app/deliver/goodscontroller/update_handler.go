package goodscontroller

import (
	"crudserver/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func (c *controller) updateHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c.log.Info("Updating good")
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

		err = c.storage.UpdateGood(ctx.Context(), good)
		if err != nil {
			c.log.Errorf("Failed to update good: %s", err)
			return fiber.NewError(fiber.StatusNotFound, c.notFoundError())
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
