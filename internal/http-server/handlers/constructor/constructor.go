package constructor

import (
	"fmt"
	"net/http"

	"constructor/components/pages"
	"constructor/internal/http-server/handlers/course"
	"constructor/internal/http-server/middlewares"
	"constructor/internal/lib/render"

	"go.uber.org/zap"
)

func ConstructorPage(logger *zap.SugaredLogger, courseSerivce course.CourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		switch r.Method {
		case http.MethodGet:
			v, ok := r.Context().Value(middlewares.UserIDKey).(string)
			if !ok {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			_ = v

			courses, err := courseSerivce.GetAllCourses(ctx)
			if err != nil {
				logger.Errorf("Failed to fetch courses: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			urls := make([]string, 0)
			for _, v := range courses {
				urls = append(urls, fmt.Sprintf("/course/%d", v.ID))
			}

			if err := render.Render(r.Context(), w, pages.ConstructorBase(pages.Constructor(courses, urls))); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
