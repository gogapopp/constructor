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

func (c *courseService) GetCourseByID(ctx context.Context, courseID int) (*model.Course, error) {
	return c.courseStore.GetCourseByID(ctx, courseID)
}

func (c *courseService) GetAllCourses(ctx context.Context) ([]model.Course, error) {
	return c.courseStore.GetAllCourses(ctx)
}
