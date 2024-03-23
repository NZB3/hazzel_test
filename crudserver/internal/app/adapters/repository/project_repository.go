package repository

import (
	"context"
	"crudserver/internal/app/models"
)

func (r *repo) CreateProject(ctx context.Context, project *models.Project) error {
	stmt := "INSERT INTO projects (name) VALUES ($1);"
	_, err := r.db.ExecContext(ctx, stmt, project.Name)
	if err != nil {
		r.log.Errorf("Failed to create project: %s", err)
		return err
	}

	return nil
}

func (r *repo) GetProjects(ctx context.Context) ([]*models.Project, error) {
	stmt := "SELECT * FROM projects;"
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		r.log.Errorf("Failed to get projects: %s", err)
		return nil, err
	}

	var projects []*models.Project
	for rows.Next() {
		project := &models.Project{}
		err = rows.Scan(&project.ID, &project.Name, &project.CreatedAt)
		if err != nil {
			r.log.Errorf("Failed to scan project: %s", err)
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (r *repo) GetProjectByID(ctx context.Context, id int) (*models.Project, error) {
	stmt := "SELECT * FROM projects WHERE id = $1;"
	row := r.db.QueryRowContext(ctx, stmt, id)

	var project *models.Project
	err := row.Scan(&project.ID, &project.Name, &project.CreatedAt)
	if err != nil {
		r.log.Errorf("Failed to get project: %s", err)
		return nil, err
	}

	return project, nil
}

func (r *repo) UpdateProject(ctx context.Context, project *models.Project) error {
	stmt := `BEGIN; 
LOCK TABLE goods IN ACCESS EXCLUSIVE MODE;
UPDATE projects SET name = $1 WHERE id = $2;
COMMIT;`
	_, err := r.db.ExecContext(ctx, stmt, project.Name, project.ID)
	if err != nil {
		r.log.Errorf("Failed to update project: %s", err)
		return err
	}

	return nil
}

func (r *repo) DeleteProject(ctx context.Context, id int) error {
	stmt := `BEGIN; 
LOCK TABLE goods IN ACCESS EXCLUSIVE MODE;
DELETE FROM projects WHERE id = $1;
COMMIT;`
	_, err := r.db.ExecContext(ctx, stmt, id)
	if err != nil {
		r.log.Errorf("Failed to delete project: %s", err)
		return err
	}

	return nil
}
