package home

import (
	"net/http"

	"constructor/internal/lib/render"

	"constructor/components"

	"go.uber.org/zap"
)

func HomePage(logger *zap.SugaredLogger) http.HandlerFunc {
	const op = "handlers.home.HomePage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = ctx
		switch r.Method {
		case http.MethodGet:
			if err := render.Render(r.Context(), w, components.HomeBase(components.Home("Constructor Constructor Constructor"))); err != nil {
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
