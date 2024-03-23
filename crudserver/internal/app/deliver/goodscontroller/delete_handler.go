package goodscontroller

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (c *controller) deleteHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c.log.Info("Deleting good")
		idQuery := ctx.Params("id")
		if idQuery == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Bad request")
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Bad request")
		}

		err = c.storage.DeleteGood(ctx.Context(), id)
		if err != nil {
			c.log.Errorf("Failed to delete good: %s", err)

			return fiber.NewError(fiber.StatusNotFound, c.notFoundError())
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
