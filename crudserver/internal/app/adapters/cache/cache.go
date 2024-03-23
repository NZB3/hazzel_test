package cache

import (
	"crudserver/internal/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

type cache struct {
	redis *redis.Client
	log   logger.Logger
}

func New(log logger.Logger) (*cache, error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	pwd := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	return &cache{redis: conn, log: log}, nil
}
