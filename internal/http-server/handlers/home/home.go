package home

import (
	"net/http"

	"constructor/components/pages"
	"constructor/internal/lib/render"

	"go.uber.org/zap"
)

func HomePage(logger *zap.SugaredLogger) http.HandlerFunc {
	const op = "handlers.home.HomePage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = ctx
		switch r.Method {
		case http.MethodGet:
			if err := render.Render(r.Context(), w, pages.HomeBase(pages.Home("Constructor"))); err != nil {
				logger.Errorf("%s: %w", op, err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
