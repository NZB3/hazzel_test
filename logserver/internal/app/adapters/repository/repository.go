package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"logserver/internal/logger"
	"os"
)

type repository struct {
	log logger.Logger
	db  *sql.DB
}

func New(log logger.Logger) (*repository, error) {
	log.Info("Creating repository")

	dsn := fmt.Sprintf(
		"clickhouse://%s:%s@%s:%s/%s",
		os.Getenv("CLICKHOUSE_USER"),
		os.Getenv("CLICKHOUSE_PASSWORD"),
		os.Getenv("CLICKHOUSE_HOST"),
		os.Getenv("CLICKHOUSE_PORT"),
		os.Getenv("CLICKHOUSE_DB"),
	)
	conn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Errorf("Failed to connect to ClickHouse: %s", err)
		return nil, err
	}

	return &repository{
		log: log,
		db:  conn,
	}, nil
}
