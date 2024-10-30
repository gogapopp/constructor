package constructor

import (
	"net/http"

	"constructor/components/pages"
	"constructor/internal/http-server/middlewares"
	"constructor/internal/lib/render"

	"go.uber.org/zap"
)

func ConstructorPage(logger *zap.SugaredLogger) http.HandlerFunc {
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
			if err := render.Render(r.Context(), w, pages.ConstructorBase(pages.Constructor())); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
