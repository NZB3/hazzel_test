package app

import (
	"context"
	"crudserver/internal/app/adapters/cache"
	"crudserver/internal/app/adapters/messagebrocker"
	"crudserver/internal/app/adapters/repository"
	"crudserver/internal/app/deliver/goodscontroller"
	"crudserver/internal/app/deliver/projectscontroller"
	"crudserver/internal/app/services/goodsservice"
	"crudserver/internal/app/services/projectservice"
	"crudserver/internal/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

type app struct {
	log    logger.Logger
	server *fiber.App
}

func New() (*app, error) {
	msgBroker, err := messagebrocker.New()
	if err != nil {
		return nil, err
	}

	log := logger.New(msgBroker.GetWriter("log_event"))
	log.SetServiceName("app")

	repo, err := repository.New(log)
	if err != nil {
		log.Errorf("Failed to init repository: %s", err)
		return nil, err
	}

	c, err := cache.New(log)
	if err != nil {
		log.Errorf("Failed to init cache: %s", err)
		return nil, err
	}

	goodService := goodsservice.New(log, c, repo, msgBroker)
	projectService := projectservice.New(log, c, repo)

	fiberApp := fiber.New()

	goodsController := goodscontroller.New(log, goodService)
	goodsController.AddRoutsToApp(fiberApp)

	projectController := projectscontroller.New(log, projectService)
	projectController.AddRoutsToApp(fiberApp)

	return &app{
		log:    log,
		server: fiberApp,
	}, nil
}

func (a *app) Run() error {
	a.log.Info("Starting server")
	port := os.Getenv("GOODS_API_PORT")
	return a.server.Listen(fmt.Sprintf(":%s", port))
}

func (a *app) Stop(complete chan<- struct{}) error {
	a.log.Info("Stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.ShutdownWithContext(ctx); err != nil {
		a.log.Errorf("Failed to shutdown server: %s", err)
		return err
	}

	close(complete)
	return nil
}
