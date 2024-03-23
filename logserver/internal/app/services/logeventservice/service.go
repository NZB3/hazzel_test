package logeventservice

import (
	"logserver/internal/app/models"
	"logserver/internal/logger"
)

type logEventStorage interface {
	SaveLogEvent(event models.LogEvent) error
}

type service struct {
	log     logger.Logger
	storage logEventStorage
	queue   chan models.LogEvent
	count   int
}

func New(log logger.Logger, storage logEventStorage) *service {
	return &service{
		log:     log,
		storage: storage,
		count:   0,
		queue:   make(chan models.LogEvent, 1),
	}
}

func (s *service) SaveLogEvent(event models.LogEvent) error {
	s.log.Info("Saving log event")
	if s.count == 5 {
		s.count = 0
		for e := range s.queue {
			if err := s.storage.SaveLogEvent(e); err != nil {
				return err
			}
		}
		return nil
	}

	s.count++
	s.log.Info("Log event added to queue")
	go func() {
		s.queue <- event
	}()
	return nil
}
