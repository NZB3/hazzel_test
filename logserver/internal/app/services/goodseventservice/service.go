package goodseventservice

import (
	"logserver/internal/app/models"
	"logserver/internal/logger"
)

type goodsEventStorage interface {
	SaveGoodsEvent(event models.GoodsEvent) error
}

type service struct {
	log      logger.Logger
	storage  goodsEventStorage
	queue    chan models.GoodsEvent
	stopChan chan struct{}
	count    int
}

func New(log logger.Logger, storage goodsEventStorage) *service {
	return &service{
		log:      log,
		storage:  storage,
		count:    0,
		stopChan: make(chan struct{}),
		queue:    make(chan models.GoodsEvent, 1),
	}
}

func (s *service) SaveGoodsEvent(event models.GoodsEvent) error {
	s.log.Info("Saving goods event")
	if s.count == 5 {
		s.count = 0
		for e := range s.queue {
			if err := s.storage.SaveGoodsEvent(e); err != nil {
				return err
			}
		}
		return nil
	}

	s.count++
	s.log.Info("Goods event added to queue")
	go func() {
		s.queue <- event
	}()
	return nil
}
