package service

import (
	"constructor/internal/model"
	"context"
)

func (c *courseService) CreateCourse(ctx context.Context, course model.Course) error {
	if err := c.validator.Struct(course); err != nil {
		return err
	}

	return c.courseStore.CreateCourse(ctx, course)
}
