package projectservice

import (
	"context"
	"crudserver/internal/app/models"
	"crudserver/internal/logger"
)

type projectStorage interface {
	CreateProject(ctx context.Context, project *models.Project) error
	GetProjects(ctx context.Context) ([]*models.Project, error)
	GetProjectByID(ctx context.Context, id int) (*models.Project, error)
	UpdateProject(ctx context.Context, project *models.Project) error
	DeleteProject(ctx context.Context, id int) error
}

type projectCache interface {
	CacheProject(ctx context.Context, project *models.Project) error
	GetCachedProject(ctx context.Context, id int) (*models.Project, error)
	DeleteCachedProject(ctx context.Context, id int) error
}

type service struct {
	storage projectStorage
	cache   projectCache
	log     logger.Logger
}

func New(log logger.Logger, cache projectCache, storage projectStorage) *service {
	return &service{
		cache:   cache,
		storage: storage,
		log:     log,
	}
}

func (s *service) GetProjects(ctx context.Context) ([]*models.Project, error) {
	s.log.Info("Get projects")
	return s.storage.GetProjects(ctx)
}

func (s *service) GetProjectByID(ctx context.Context, id int) (*models.Project, error) {
	s.log.Info("Get project by id")
	project, err := s.cache.GetCachedProject(ctx, id)
	if err != nil {
		return nil, err
	}

	if project == nil {
		project, err = s.storage.GetProjectByID(ctx, id)
		if err != nil {
			return nil, err
		}
		err = s.cache.CacheProject(ctx, project)
		if err != nil {
			return nil, err
		}
	}

	return project, nil
}

func (s *service) SaveProject(ctx context.Context, project *models.Project) error {
	s.log.Info("Create project")
	return s.storage.CreateProject(ctx, project)
}

func (s *service) UpdateProject(ctx context.Context, project *models.Project) error {
	s.log.Info("Update project")
	if err := s.storage.UpdateProject(ctx, project); err != nil {
		return err
	}

	if err := s.cache.CacheProject(ctx, project); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteProject(ctx context.Context, id int) error {
	s.log.Info("Delete project")
	if err := s.storage.DeleteProject(ctx, id); err != nil {
		return err
	}

	if err := s.cache.DeleteCachedProject(ctx, id); err != nil {
		return err
	}

	return nil
}
