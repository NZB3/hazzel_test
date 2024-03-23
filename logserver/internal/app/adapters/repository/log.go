package repository

import "logserver/internal/app/models"

func (r *repository) SaveLogEvent(log models.LogEvent) error {
	stmt := `INSERT INTO log_event (time, msg, level) VALUES (?, ?, ?)`
	_, err := r.db.Exec(stmt, log.Time, log.Msg, log.Level)
	if err != nil {
		return err
	}

	return nil
}
