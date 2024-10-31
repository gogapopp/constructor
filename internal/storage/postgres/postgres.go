package postgres

import (
	"constructor/sqlembed"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dsn string) (*pgxpool.Pool, error) {
	const op = "postgres.New"
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer func() {
		err = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, sqlembed.UsersTable)
	_, err = tx.Exec(ctx, sqlembed.CoursesTable)
	_, err = tx.Exec(ctx, sqlembed.CourseModulesTable)
	_, err = tx.Exec(ctx, sqlembed.CourseAccessTable)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return conn, tx.Commit(ctx)
}
