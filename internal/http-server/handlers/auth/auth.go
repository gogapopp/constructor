package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"constructor/components/pages"
	"constructor/internal/lib/render"
	"constructor/internal/model"
	"constructor/internal/service"
	"constructor/internal/storage"

	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

var valErr validator.ValidationErrors

type AuthService interface {
	SignUp(ctx context.Context, user model.SignUpUser) error
	SignIn(ctx context.Context, user model.SignInUser) (string, error)
}

func SignUpPage(logger *zap.SugaredLogger, a AuthService) http.HandlerFunc {
	const op = "handlers.auth.SignUpPage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		switch r.Method {
		case http.MethodPost:
			userName := r.FormValue("username")
			userEmail := r.FormValue("email")
			userPassword := r.FormValue("password")
			userPasswordConfirm := r.FormValue("password_confirm")

			if userPassword != userPasswordConfirm {
				if err := render.Render(ctx, w, pages.SignUpBase(pages.SignUp("Passwords doesn't equals"))); err != nil {
					logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}
			}

			err := a.SignUp(ctx, model.SignUpUser{
				Username:     userName,
				Email:        userEmail,
				PasswordHash: userPassword,
				CreatedAt:    time.Now(),
				Role:         "user",
			})
			if err != nil {
				logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
				var errMsg string
				if errors.Is(err, storage.ErrUserExists) {
					errMsg = "User already exists"
				} else if errors.As(err, &valErr) {
					errMsg = "Email, password and username is required field"
				} else if errors.Is(err, service.ErrUndefinedRole) {
					errMsg = "Undefined role (available roles: admin, user)"
				} else {
					errMsg = "Something went wrong"
				}
				if err := render.Render(ctx, w, pages.SignUpBase(pages.SignUp(errMsg))); err != nil {
					logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}
			}

			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return

		case http.MethodGet:
			var errMsg string
			if r.URL.Query().Get("redirected") == "true" {
				errMsg = "You need to signin or signup"
			}
			if err := render.Render(ctx, w, pages.SignUpBase(pages.SignUp(errMsg))); err != nil {
				logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func SignInPage(logger *zap.SugaredLogger, a AuthService) http.HandlerFunc {
	const op = "handlers.auth.SignInPage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		switch r.Method {
		case http.MethodPost:
			emailOrUsername := r.FormValue("email_or_username")
			userPassword := r.FormValue("password")

			token, err := a.SignIn(ctx, model.SignInUser{
				Email:        emailOrUsername,
				Username:     emailOrUsername,
				PasswordHash: userPassword,
			})
			if err != nil {
				logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
				var errMsg string
				if errors.Is(err, storage.ErrUserNotExist) {
					errMsg = "Invalid email/username or password"
				} else if errors.As(err, &valErr) {
					errMsg = "Email or username and password is required field"
				} else {
					errMsg = "Something went wrong"
				}
				if err := render.Render(ctx, w, pages.SignInBase(pages.SignIn(errMsg))); err != nil {
					logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}
			}

			cookie := &http.Cookie{
				Name:     "ssid",
				Value:    token,
				Path:     "/",
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				Secure:   true}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/constructor", http.StatusSeeOther)
			return

		case http.MethodGet:
			var errMsg string
			if r.URL.Query().Get("redirected") == "true" {
				errMsg = "You need to signin or signup"
			}
			if err := render.Render(ctx, w, pages.SignInBase(pages.SignIn(errMsg))); err != nil {
				logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
