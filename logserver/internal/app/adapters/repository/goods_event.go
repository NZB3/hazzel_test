package repository

import "logserver/internal/app/models"

func (r *repository) SaveGoodsEvent(event models.GoodsEvent) error {
	stmt := `INSERT INTO goods_event (id, project_id, name, description, priority, removed, event_time) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(
		stmt,
		event.ID,
		event.ProjectID,
		event.Name,
		event.Description,
		event.Priority,
		event.Removed,
		event.EventTime,
	)

	if err != nil {
		r.log.Errorf("Failed to save goods event: %s", err)
		return err
	}

	return nil
}
