package sqlite

import (
	"constructor/sqlembed"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func New(dsn string) (*sql.DB, error) {
	const op = "sqlite.New"

	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tx, err := conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer func() {
		err = tx.Rollback()
	}()

	_, err = tx.Exec(sqlembed.UsersTable)
	_, err = tx.Exec(sqlembed.CoursesTable)
	_, err = tx.Exec(sqlembed.CourseModulesTable)
	_, err = tx.Exec(sqlembed.CourseAccessTable)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return conn, tx.Commit()
}
