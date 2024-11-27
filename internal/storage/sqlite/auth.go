package sqlite

import (
	"constructor/internal/model"
	"constructor/internal/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type authStorage struct {
	conn *sql.DB
}

func NewAuthStorage(conn *sql.DB) *authStorage {
	return &authStorage{
		conn: conn,
	}
}

func (s *authStorage) SignUp(ctx context.Context, user model.SignUpUser) error {
	const (
		op    = "sqlite.SignUp"
		query = "INSERT INTO users (email, username, password_hash, created_at, role) VALUES (?, ?, ?, ?, ?);" // TODO: add metainfo
	)

	_, err := s.conn.Exec(query, user.Email, user.Username, user.PasswordHash, user.CreatedAt, user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *authStorage) SignIn(ctx context.Context, user model.SignInUser) (int, string, error) {
	const (
		op    = "sqlite.SignIn"
		query = "SELECT user_id, role FROM users WHERE username=? AND password_hash=?"
	)
	var (
		userID int
		role   string
	)

	row := s.conn.QueryRow(query, user.Username, user.PasswordHash)

	err := row.Scan(&userID, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", fmt.Errorf("%s: %w", op, storage.ErrUserNotExist)
		}

		return 0, "", fmt.Errorf("%s: %w", op, err)
	}

	return userID, role, nil
}
