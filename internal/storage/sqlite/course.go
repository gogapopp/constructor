package sqlite

import (
	"constructor/internal/model"
	"context"
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

func (c *courseStorage) CreateCourse(ctx context.Context, course model.Course) error {
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `INSERT INTO courses (title, description, creator_id, created_at, difficulty_level, cover_image) VALUES (?, ?, ?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, query, course.Title, course.Description, course.CreatorID, course.CreatedAt, course.DifficultyLevel, course.CoverImage)
	if err != nil {
		return fmt.Errorf("failed to create course: %w", err)
	}

	courseID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	for _, module := range course.Modules {
		moduleQuery := `INSERT INTO course_modules (course_id, title, description) VALUES (?, ?, ?)`
		res, err := tx.ExecContext(ctx, moduleQuery, courseID, module.Title, module.Description)
		if err != nil {
			return fmt.Errorf("failed to create module: %w", err)
		}

		moduleID, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last insert id for module: %w", err)
		}

		for _, lesson := range module.Lessons {
			lessonQuery := `INSERT INTO course_lessons (module_id, title, type, content, video_url, resource_url) VALUES (?, ?, ?, ?, ?, ?)`
			_, err := tx.ExecContext(ctx, lessonQuery, moduleID, lesson.Title, lesson.Type, lesson.Content, lesson.VideoURL, lesson.ResourceURL)
			if err != nil {
				return fmt.Errorf("failed to create lesson: %w", err)
			}
		}
	}

	return tx.Commit()
}
