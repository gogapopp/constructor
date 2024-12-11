package sqlite

import (
	"constructor/internal/model"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
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

	query := `INSERT INTO courses (title, description, creator_id, created_at, difficulty_level) VALUES (?, ?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, query, course.Title, course.Description, course.CreatorID, course.CreatedAt, course.DifficultyLevel)
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

func (c *courseStorage) GetCourseByID(ctx context.Context, courseID int) (*model.Course, error) {
	// Retrieve course details
	courseQuery := `
        SELECT 
            course_id, 
            title, 
            description, 
            creator_id, 
            created_at, 
            difficulty_level
        FROM courses 
        WHERE course_id = ?
    `
	var course model.Course
	var createdAtStr string

	err := c.conn.QueryRowContext(ctx, courseQuery, courseID).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.CreatorID,
		&createdAtStr,
		&course.DifficultyLevel,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("course not found")
		}
		return nil, fmt.Errorf("query course: %w", err)
	}

	// Parse created_at timestamp
	course.CreatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("parse created_at: %w", err)
	}

	// Retrieve modules
	modulesQuery := `
        SELECT 
            module_id,
            title, 
            description
        FROM course_modules 
        WHERE course_id = ?
        ORDER BY order_index
    `
	rows, err := c.conn.QueryContext(ctx, modulesQuery, courseID)
	if err != nil {
		return nil, fmt.Errorf("query modules: %w", err)
	}
	defer rows.Close()

	var modules []model.Module
	var moduleIDs []int
	for rows.Next() {
		var module model.Module
		var moduleID int
		if err := rows.Scan(&moduleID, &module.Title, &module.Description); err != nil {
			return nil, fmt.Errorf("scan module: %w", err)
		}
		module.ModuleID = moduleID
		modules = append(modules, module)
		moduleIDs = append(moduleIDs, moduleID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	// Retrieve lessons for all modules
	if len(moduleIDs) == 0 {
		course.Modules = modules
		return &course, nil
	}

	// Create placeholders
	placeholderSlice := make([]string, len(moduleIDs))
	for i := range placeholderSlice {
		placeholderSlice[i] = "?"
	}
	placeholders := strings.Join(placeholderSlice, ",")

	// Create lessons query
	lessonsQuery := fmt.Sprintf(`
    SELECT 
        lesson_id,
        module_id,
        title,
        type,
        content,
        video_url,
        resource_url
    FROM course_lessons
    WHERE module_id IN (%s)
`, placeholders)

	// Prepare arguments for query
	args := make([]any, len(moduleIDs))
	for i, id := range moduleIDs {
		args[i] = id
	}

	// Execute query
	lessonRows, err := c.conn.QueryContext(ctx, lessonsQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("query lessons: %w", err)
	}
	defer lessonRows.Close()

	lessonsMap := make(map[int][]model.Lesson)
	for lessonRows.Next() {
		var lesson model.Lesson
		var moduleID int
		if err := lessonRows.Scan(&lesson.LessonID, &moduleID, &lesson.Title, &lesson.Type, &lesson.Content, &lesson.VideoURL, &lesson.ResourceURL); err != nil {
			return nil, fmt.Errorf("scan lesson: %w", err)
		}
		lessonsMap[moduleID] = append(lessonsMap[moduleID], lesson)
	}

	if err := lessonRows.Err(); err != nil {
		return nil, fmt.Errorf("lesson rows error: %w", err)
	}

	// Assign lessons to modules
	for i, module := range modules {
		if lessons, ok := lessonsMap[module.ModuleID]; ok {
			modules[i].Lessons = lessons
		}
	}

	course.Modules = modules
	return &course, nil
}

func (c *courseStorage) GetAllCourses(ctx context.Context) ([]model.Course, error) {
	query := `
		SELECT
			course_id,
			title,
			description,
			creator_id,
			created_at,
			difficulty_level
		FROM courses
	`
	rows, err := c.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query courses: %w", err)
	}
	defer rows.Close()

	var courses []model.Course
	for rows.Next() {
		var course model.Course
		var createdAtStr string
		if err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.CreatorID,
			&createdAtStr,
			&course.DifficultyLevel,
		); err != nil {
			return nil, fmt.Errorf("scan course: %w", err)
		}

		course.CreatedAt, err = time.Parse("2006-01-02 15:04:05.9999999-07:00", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("parse created_at: %w", err)
		}

		modules, err := c.getModulesByCourseID(ctx, course.ID)
		if err != nil {
			return nil, fmt.Errorf("get modules: %w", err)
		}
		course.Modules = modules

		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return courses, nil
}

func (c *courseStorage) getModulesByCourseID(ctx context.Context, courseID int) ([]model.Module, error) {
	query := `
		SELECT
			module_id,
			title,
			description
		FROM course_modules
		WHERE course_id = ?
		ORDER BY order_index
	`
	rows, err := c.conn.QueryContext(ctx, query, courseID)
	if err != nil {
		return nil, fmt.Errorf("query modules: %w", err)
	}
	defer rows.Close()

	var modules []model.Module
	for rows.Next() {
		var module model.Module
		var moduleID int
		if err := rows.Scan(&moduleID, &module.Title, &module.Description); err != nil {
			return nil, fmt.Errorf("scan module: %w", err)
		}

		// Retrieve lessons for each module (if needed)
		// Note: This is a placeholder. You'd need to create a lessons table
		// and implement lesson retrieval similar to module retrieval
		module.Lessons = []model.Lesson{} // Placeholder for lessons

		modules = append(modules, module)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return modules, nil
}
