package cache

import (
	"context"
	"crudserver/internal/app/models"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func (c *cache) GetCachedGood(ctx context.Context, id int) (*models.Good, error) {
	var good *models.Good

	err := c.redis.Get(ctx, strconv.Itoa(id)).Scan(&good)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		c.log.Errorf("Failed to get cached good: %s", err)
		return nil, err
	}

	return good, nil
}

func (c *cache) CacheGood(ctx context.Context, good *models.Good) error {
	err := c.redis.Set(ctx, strconv.Itoa(good.ID), good, time.Minute).Err()
	if err != nil {
		c.log.Errorf("Failed to set good in cache: %s", err)
		return err
	}

	return nil
}

func (c *cache) DeleteCachedGood(ctx context.Context, id int) error {
	err := c.redis.Del(ctx, strconv.Itoa(id)).Err()
	if err != nil {
		c.log.Errorf("Failed to delete good from cache: %s", err)
		return err
	}

	return nil
}
