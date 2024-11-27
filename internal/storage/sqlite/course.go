package sqlite

import (
	"database/sql"
	"fmt"
)

type courseStorage struct {
	conn *sql.DB
}

func NewCourseStorage(conn *sql.DB) *courseStorage {
	return &courseStorage{
		conn: conn,
	}
}

func (c *courseStorage) CheckAndExpireCourseAccess() error {
	const (
		op    = "sqlite.CheckAndExpireCourseAccess"
		query = "UPDATE course_access SET is_active = FALSE, access_status = 'expired' WHERE valid_until < datetime('now') AND access_status = 'active'"
	)

	_, err := c.conn.Exec(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
