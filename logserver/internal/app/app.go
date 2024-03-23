package app

import (
	"context"
	"logserver/internal/app/adapters/repository"
	"logserver/internal/app/deliver/eventlistener"
	"logserver/internal/app/deliver/goodscontroller"
	"logserver/internal/app/deliver/logcontroller"
	"logserver/internal/app/services/goodseventservice"
	"logserver/internal/app/services/logeventservice"
	"logserver/internal/logger"
)

type listener interface {
	Listen(ctx context.Context) error
	Stop()
}

type logService struct {
	log logger.Logger
	el  listener
}

func New() (*logService, error) {
	log := logger.New()
	log.Info("Logger sat up")

	repo, err := repository.New(log)
	if err != nil {
		log.Errorf("Failed to create repository: %s", err)
		return nil, err
	}
	log.Info("Repository sat up")

	goodsEventService := goodseventservice.New(log, repo)
	logEventService := logeventservice.New(log, repo)

	goodsController := goodscontroller.New(log, goodsEventService)
	logEventController := logcontroller.New(log, logEventService)

	log.Info("Controllers sat up")

	el, err := eventlistener.New(log)
	if err != nil {
		log.Errorf("Failed to create event listener: %s", err)
		return nil, err
	}

	log.Info("Event listener sat up")

	el.HandleFunc("goods_event", goodsController.GoodsHandler)
	el.HandleFunc("log_event", logEventController.LogHandler)

	log.Info("Event handlers sat up")

	return &logService{
		log: log,
		el:  el,
	}, err
}

func (l *logService) Run(ctx context.Context) error {
	l.log.Info("Starting log services")
	if err := l.el.Listen(ctx); err != nil {
		l.log.Errorf("Failed to start event listener: %s", err)
		return err
	}

	return nil
}

func (l *logService) Stop(complete chan<- struct{}) error {
	l.log.Info("Stopping log services")
	l.el.Stop()
	close(complete)
	return nil
}
