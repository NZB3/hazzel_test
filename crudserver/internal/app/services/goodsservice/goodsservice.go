package goodsservice

import (
	"context"
	"crudserver/internal/app/models"
	"crudserver/internal/logger"
)

type goodCache interface {
	CacheGood(ctx context.Context, good *models.Good) error
	GetCachedGood(ctx context.Context, id int) (*models.Good, error)
	DeleteCachedGood(ctx context.Context, id int) error
}

type eventBus interface {
	SendGoodEvent(good *models.Good) error
}

type goodStorage interface {
	CreateGood(ctx context.Context, good *models.Good) error
	GetGoods(ctx context.Context) ([]*models.Good, error)
	GetGoodByID(ctx context.Context, id int) (*models.Good, error)
	UpdateGood(ctx context.Context, good *models.Good) error
	DeleteGood(ctx context.Context, id int) error
}

type service struct {
	cache    goodCache
	storage  goodStorage
	eventBus eventBus
	log      logger.Logger
}

func New(log logger.Logger, cache goodCache, storage goodStorage, eventBus eventBus) *service {
	return &service{
		cache:    cache,
		storage:  storage,
		eventBus: eventBus,
		log:      log,
	}
}

func (s *service) GetGoods(ctx context.Context) ([]*models.Good, error) {
	s.log.Info("Get goods")
	return s.storage.GetGoods(ctx)
}

func (s *service) GetGoodByID(ctx context.Context, id int) (*models.Good, error) {
	s.log.Info("Get good by id")
	good, err := s.cache.GetCachedGood(ctx, id)
	if err != nil {
		return nil, err
	}

	if good == nil {
		good, err = s.storage.GetGoodByID(ctx, id)
		if err != nil {
			return nil, err
		}
		err = s.cache.CacheGood(ctx, good)
		if err != nil {
			return nil, err
		}
	}

	return good, nil
}

func (s *service) SaveGood(ctx context.Context, good *models.Good) error {
	s.log.Info("Create good")

	if err := s.storage.CreateGood(ctx, good); err != nil {
		return err
	}

	if err := s.eventBus.SendGoodEvent(good); err != nil {
		s.log.Errorf("Failed to send good event: %s", err)
	}

	return nil
}

func (s *service) UpdateGood(ctx context.Context, good *models.Good) error {
	s.log.Info("Update good")
	if err := s.storage.UpdateGood(ctx, good); err != nil {
		return err
	}

	if err := s.cache.CacheGood(ctx, good); err != nil {
		return err
	}

	if err := s.eventBus.SendGoodEvent(good); err != nil {
		s.log.Errorf("Failed to send good event: %s", err)
	}

	return nil
}

func (s *service) DeleteGood(ctx context.Context, id int) error {
	s.log.Info("Delete good")
	if err := s.storage.DeleteGood(ctx, id); err != nil {
		return err
	}

	if err := s.cache.DeleteCachedGood(ctx, id); err != nil {
		return err
	}

	if err := s.eventBus.SendGoodEvent(&models.Good{ID: id, Removed: true}); err != nil {
		s.log.Errorf("Failed to send good event: %s", err)
	}

	return nil
}
