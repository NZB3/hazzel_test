package goodscontroller

import (
	"crudserver/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"strconv"
)

func (c *controller) getOneHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c.log.Info("Getting good")
		idQuery := ctx.Params("id")
		if idQuery == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Bad request")
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Bad request")
		}

		good, err := c.storage.GetGoodByID(ctx.Context(), id)
		if err != nil {
			c.log.Errorf("Failed to get good: %s", err)
			return fiber.NewError(fiber.StatusInternalServerError, c.internalServerError())
		}

		return ctx.JSON(&goodsResponse{
			ProjectID:   good.ProjectID,
			Name:        good.Name,
			Description: good.Description,
			Priority:    good.Priority,
			CreatedAt:   good.CreatedAt,
		})
	}
}

func (c *controller) getAllHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c.log.Info("Getting goods")
		goods, err := c.storage.GetGoods(ctx.Context())
		if err != nil {
			c.log.Errorf("Failed to get goods: %s", err)
			return fiber.NewError(fiber.StatusInternalServerError, c.internalServerError())
		}

		return ctx.JSON(lo.Map(goods, func(good *models.Good, _ int) goodsResponse {
			return goodsResponse{
				ProjectID:   good.ProjectID,
				Name:        good.Name,
				Description: good.Description,
				Priority:    good.Priority,
				CreatedAt:   good.CreatedAt,
			}
		}))
	}
}
