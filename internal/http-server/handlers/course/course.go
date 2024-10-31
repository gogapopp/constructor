package course

import (
	"constructor/components/pages"
	"constructor/internal/http-server/middlewares"
	"constructor/internal/lib/render"
	"net/http"

	"go.uber.org/zap"
)

func CoursePage(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = ctx
		switch r.Method {
		case http.MethodGet:
			v, ok := r.Context().Value(middlewares.UserIDKey).(string)
			if !ok {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			_ = v
			// TODO:
			if err := render.Render(r.Context(), w, pages.CourseBase(pages.Course())); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
