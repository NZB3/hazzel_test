package repository

import (
	"context"
	"crudserver/internal/app/models"
)

func (r *repo) GetGoods(ctx context.Context) ([]*models.Good, error) {
	stmt := "SELECT * FROM goods WHERE removed = false;"
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		r.log.Errorf("Failed to get goods: %s", err)
		return nil, err
	}

	var goods []*models.Good
	for rows.Next() {
		good := &models.Good{}
		err = rows.Scan(&good.ID, &good.ProjectID, &good.Name, &good.Description, &good.Priority, &good.Removed, &good.CreatedAt)
		if err != nil {
			r.log.Errorf("Failed to scan good: %s", err)
			return nil, err
		}
		goods = append(goods, good)
	}

	return goods, nil
}

func (r *repo) GetGoodByID(ctx context.Context, id int) (*models.Good, error) {
	stmt := "SELECT * FROM goods WHERE id = $1;"
	row := r.db.QueryRowContext(ctx, stmt, id)

	good := &models.Good{}
	err := row.Scan(&good.ID, &good.ProjectID, &good.Name, &good.Description, &good.Priority, &good.Removed, &good.CreatedAt)
	if err != nil {
		r.log.Errorf("Failed to get good: %s", err)
		return nil, err
	}

	return good, nil
}

func (r *repo) GetGoodByIDAndProjectID(ctx context.Context, id, projectID int) (*models.Good, error) {
	stmt := "SELECT * FROM goods WHERE id = $1 AND project_id = $2;"
	row := r.db.QueryRowContext(ctx, stmt, id, projectID)

	var good *models.Good
	err := row.Scan(&good.ID, &good.ProjectID, &good.Name, &good.Description, &good.Priority, &good.Removed, &good.CreatedAt)
	if err != nil {
		r.log.Errorf("Failed to get good: %s", err)
		return nil, err
	}

	return good, nil
}

func (r *repo) CreateGood(ctx context.Context, good *models.Good) error {
	stmt := `INSERT INTO goods (project_id, name, description, priority) VALUES ($1, $2, $3, $4);`
	_, err := r.db.ExecContext(ctx, stmt, good.ProjectID, good.Name, good.Description, good.Priority)
	if err != nil {
		r.log.Errorf("Failed to create good: %s", err)
		return err
	}

	return nil
}

func (r *repo) UpdateGood(ctx context.Context, good *models.Good) error {
	stmt := `BEGIN; 
LOCK TABLE goods IN ACCESS EXCLUSIVE MODE; 
UPDATE goods SET name = $1, description = $2, priority = $3 WHERE id = $4; 
COMMIT;`
	_, err := r.db.ExecContext(ctx, stmt, good.Name, good.Description, good.Priority, good.ID)
	if err != nil {
		r.log.Errorf("Failed to update good: %s", err)
		return err
	}

	return nil
}

func (r *repo) DeleteGood(ctx context.Context, id int) error {
	stmt := `BEGIN; 
LOCK TABLE goods IN ACCESS EXCLUSIVE MODE; 
UPDATE goods SET removed = true WHERE id = $1;
COMMIT;`
	_, err := r.db.ExecContext(ctx, stmt, id)
	if err != nil {
		r.log.Errorf("Failed to delete good: %s", err)
		return err
	}

	return nil
}
