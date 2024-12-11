package course

import (
	"constructor/components/pages"
	"constructor/internal/lib/render"
	"constructor/internal/model"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type CourseService interface {
	CreateCourse(ctx context.Context, course model.Course) error
	GetCourseByID(ctx context.Context, courseID int) (*model.Course, error)
	GetAllCourses(ctx context.Context) ([]model.Course, error)
}

func CoursePage(logger *zap.SugaredLogger, courseService CourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			courseIDStr := chi.URLParam(r, "id")
			courseID, err := strconv.Atoi(courseIDStr)
			if err != nil {
				http.Error(w, "invalid course ID", http.StatusBadRequest)
				return
			}
			course, err := courseService.GetCourseByID(ctx, courseID)
			if err != nil {
				logger.Errorf("failed to fetch course: %v", err)
				http.Error(w, "failed to retrieve course", http.StatusInternalServerError)
				return
			}

			logger.Infoln(course.Title, course.Description, course.DifficultyLevel, course.CreatedAt, course.CreatorID, course.Modules)
			// logger.Infoln(course.Modules[1].Description, course.Modules[1].Lessons, course.Modules[1].Title)

			if err := render.Render(ctx, w, pages.CourseBase(pages.Course(course))); err != nil {
				logger.Errorf("Failed to render course page: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		case http.MethodPost:
			http.Error(w, "not implemented", http.StatusNotImplemented)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
