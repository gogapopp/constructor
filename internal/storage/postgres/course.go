package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type courseStorage struct {
	conn *pgxpool.Pool
}

func NewCourseStorage(conn *pgxpool.Pool) *courseStorage {
	return &courseStorage{
		conn: conn,
	}
}

func (c *courseStorage) CheckAndExpireCourseAccess() error {
	const (
		op    = "postgres.CheckAndExpireCourseAccess"
		query = "UPDATE course_access SET is_active = FALSE, access_status = 'expired' WHERE valid_until < NOW() AND access_status = 'active'"
	)

	_, err := c.conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
