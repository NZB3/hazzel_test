package repository

import (
	"crudserver/internal/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type repo struct {
	db  *sql.DB
	log logger.Logger
}

func New(log logger.Logger) (*repo, error) {
	log.Info("Creating repository")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Errorf("Failed to connect to Postgres: %s", err)
		return nil, err
	}

	log.Info("Connected to Postgres")

	return &repo{db: db, log: log}, nil
}
