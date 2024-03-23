package messagebrocker

import (
	"crudserver/internal/app/models"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"io"
	"os"
	"time"
)

type messageBroker struct {
	nc *nats.Conn
}

type goodsEvent struct {
	ID          int       `json:"id,omitempty"`
	ProjectID   int       `json:"project_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Removed     bool      `json:"removed,omitempty"`
	EventTime   time.Time `json:"event_time"`
}

func (m *messageBroker) SendGoodEvent(good *models.Good) error {
	data, err := json.Marshal(&goodsEvent{
		ID:          good.ID,
		ProjectID:   good.ProjectID,
		Name:        good.Name,
		Description: good.Description,
		Priority:    good.Priority,
		Removed:     good.Removed,
		EventTime:   time.Now(),
	})
	if err != nil {
		return err
	}
	return m.nc.Publish("goods_event", data)
}

func New() (*messageBroker, error) {
	url := fmt.Sprintf("nats://%s:%s", os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"))

	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &messageBroker{nc: nc}, nil
}

func (m *messageBroker) GetWriter(topic string) io.Writer {
	return &writer{messageBroker: m, topic: topic}
}

func (m *messageBroker) Publish(topic string, p []byte) error {
	return m.nc.Publish(topic, p)
}

type writer struct {
	messageBroker *messageBroker
	topic         string
}

func (w *writer) Write(p []byte) (n int, err error) {
	err = w.messageBroker.Publish(w.topic, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
