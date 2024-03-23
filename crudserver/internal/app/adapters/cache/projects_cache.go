package cache

import (
	"context"
	"crudserver/internal/app/models"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func (c *cache) CacheProject(ctx context.Context, project *models.Project) error {
	err := c.redis.Set(ctx, strconv.Itoa(project.ID), project, time.Minute).Err()
	if err != nil {
		c.log.Errorf("Failed to set project in cache: %s", err)
		return err
	}

	return nil
}

func (c *cache) GetCachedProject(ctx context.Context, id int) (*models.Project, error) {
	var project *models.Project

	err := c.redis.Get(ctx, strconv.Itoa(id)).Scan(&project)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		c.log.Errorf("Failed to get cached project: %s", err)
		return nil, err
	}

	return project, nil
}

func (c *cache) DeleteCachedProject(ctx context.Context, id int) error {
	err := c.redis.Del(ctx, strconv.Itoa(id)).Err()
	if err != nil {
		c.log.Errorf("Failed to delete project from cache: %s", err)
		return err
	}

	return nil
}
