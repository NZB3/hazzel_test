package logcontroller

import (
	"bytes"
	"encoding/json"
	"logserver/internal/app/models"
	"logserver/internal/logger"
)

type logStorage interface {
	SaveLogEvent(log models.LogEvent) error
}

type controller struct {
	log        logger.Logger
	logStorage logStorage
}

func New(log logger.Logger, logStorage logStorage) *controller {
	return &controller{
		log:        log,
		logStorage: logStorage,
	}
}

func (c *controller) LogHandler(event models.Event) {
	var logEvent models.LogEvent
	if err := json.NewDecoder(bytes.NewReader(event.Data)).Decode(&logEvent); err != nil {
		c.log.Errorf("Failed to decode JSON: %s", err)
		return
	}
	c.log.Info("Got log event")
	if err := c.logStorage.SaveLogEvent(logEvent); err != nil {
		c.log.Errorf("Failed to save log event: %s", err)
		return
	}
}
