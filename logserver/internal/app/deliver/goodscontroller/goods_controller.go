package goodscontroller

import (
	"bytes"
	"encoding/json"
	"logserver/internal/app/models"
	"logserver/internal/logger"
)

type goodsEventStorage interface {
	SaveGoodsEvent(event models.GoodsEvent) error
}

type controller struct {
	log logger.Logger
	ges goodsEventStorage
}

func New(log logger.Logger, ges goodsEventStorage) *controller {
	return &controller{
		log: log,
		ges: ges,
	}
}

func (c *controller) GoodsHandler(event models.Event) {
	c.log.Info("Got goods event")
	var goodsEvent models.GoodsEvent
	if err := json.NewDecoder(bytes.NewReader(event.Data)).Decode(&goodsEvent); err != nil {
		c.log.Errorf("Failed to decode JSON: %s", err)
		return
	}
	c.log.Info("Got app event")
	if err := c.ges.SaveGoodsEvent(goodsEvent); err != nil {
		c.log.Errorf("Failed to save app event: %s", err)
		return
	}

	return
}
