package eventlistener

import (
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"logserver/internal/app/models"
	"logserver/internal/logger"
	"os"
)

type eventListener struct {
	log    logger.Logger
	nc     *nats.Conn
	router map[string]models.EventHandler
}

func New(log logger.Logger) (*eventListener, error) {
	log.Info("Event listener creating")

	url := fmt.Sprintf("nats://%s:%s", os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"))
	nc, err := nats.Connect(url)
	if err != nil {
		log.Errorf("Failed to connect to NATS: %s", err)
		return nil, err
	}
	log.Info("Connected to NATS")

	router := make(map[string]models.EventHandler)

	return &eventListener{
		log:    log,
		nc:     nc,
		router: router,
	}, nil
}

func (e *eventListener) HandleFunc(topic string, handler models.EventHandler) {
	e.router[topic] = handler
}

func (e *eventListener) Listen(ctx context.Context) error {
	for topic := range e.router {
		go func(topic string, handler models.EventHandler) {
			err := e.subscribe(ctx, topic, handler)
			if err != nil {
				e.log.Errorf("Failed to subscribe to topic %s: %s", topic, err)
			}
		}(topic, e.router[topic])
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (e *eventListener) Stop() {
	e.nc.Close()
}

func (e *eventListener) subscribe(ctx context.Context, topic string, handler models.EventHandler) error {
	sub, err := e.nc.SubscribeSync(topic)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := sub.NextMsgWithContext(ctx)
			if err != nil {
				if errors.Is(err, nats.ErrTimeout) {
					e.log.Info("Timeout")
					continue
				}
				return err
			}
			e.log.Info("Got message from NATS")
			event := models.Event{
				Subject: msg.Subject,
				Data:    msg.Data,
			}
			handler(event)
		}
	}
}
