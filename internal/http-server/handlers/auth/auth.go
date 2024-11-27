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

	"github.com/a-h/templ"
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
			handlePostSignUp(ctx, w, r, logger, a, op)
		case http.MethodGet:
			handleGetSignUp(ctx, w, r, logger, op)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func SignInPage(logger *zap.SugaredLogger, a AuthService) http.HandlerFunc {
	const op = "handlers.auth.SignInPage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		switch r.Method {
		case http.MethodPost:
			handlePostSignIn(ctx, w, r, logger, a, op)
		case http.MethodGet:
			handleGetSignIn(ctx, w, r, logger, op)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handlePostSignUp(ctx context.Context, w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, a AuthService, op string) {
	userName := r.FormValue("username")
	userEmail := r.FormValue("email")
	userPassword := r.FormValue("password")
	userPasswordConfirm := r.FormValue("password_confirm")

	if userPassword != userPasswordConfirm {
		renderError(ctx, w, logger, op, "Passwords don't match", pages.SignUpBase, pages.SignUp)
		return
	}

	err := a.SignUp(ctx, model.SignUpUser{
		Username:     userName,
		Email:        userEmail,
		PasswordHash: userPassword,
		CreatedAt:    time.Now(),
		Role:         "user",
	})
	if err != nil {
		handleSignUpError(ctx, w, logger, op, err)
		return
	}

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

func handlePostSignIn(ctx context.Context, w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, a AuthService, op string) {
	emailOrUsername := r.FormValue("email_or_username")
	userPassword := r.FormValue("password")

	token, err := a.SignIn(ctx, model.SignInUser{
		Email:        emailOrUsername,
		Username:     emailOrUsername,
		PasswordHash: userPassword,
	})
	if err != nil {
		handleSignInError(ctx, w, logger, op, err)
		return
	}

	cookie := &http.Cookie{
		Name:     "ssid",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/constructor", http.StatusSeeOther)
}

func handleGetSignUp(ctx context.Context, w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, op string) {
	handleGetAuth(ctx, w, r, logger, op, pages.SignUpBase, pages.SignUp)
}

func handleGetSignIn(ctx context.Context, w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, op string) {
	handleGetAuth(ctx, w, r, logger, op, pages.SignInBase, pages.SignIn)
}

func handleGetAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, op string, basePage func(templ.Component) templ.Component, authPage func(string) templ.Component) {
	var errMsg string
	if r.URL.Query().Get("redirected") == "true" {
		errMsg = "You need to sign in or sign up"
	}
	if err := render.Render(ctx, w, basePage(authPage(errMsg))); err != nil {
		renderError(ctx, w, logger, op, "internal server error", basePage, authPage)
	}
}

func handleSignUpError(ctx context.Context, w http.ResponseWriter, logger *zap.SugaredLogger, op string, err error) {
	logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
	var errMsg string
	switch {
	case errors.Is(err, storage.ErrUserExists):
		errMsg = "User already exists"
	case errors.As(err, &valErr):
		errMsg = "Email, password, and username are required fields"
	case errors.Is(err, service.ErrUndefinedRole):
		errMsg = "Undefined role (available roles: admin, user)"
	default:
		errMsg = "Something went wrong"
	}
	renderError(ctx, w, logger, op, errMsg, pages.SignUpBase, pages.SignUp)
}

func handleSignInError(ctx context.Context, w http.ResponseWriter, logger *zap.SugaredLogger, op string, err error) {
	logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
	var errMsg string
	switch {
	case errors.Is(err, storage.ErrUserNotExist):
		errMsg = "Invalid email/username or password"
	case errors.As(err, &valErr):
		errMsg = "Email or username and password are required fields"
	default:
		errMsg = "Something went wrong"
	}
	renderError(ctx, w, logger, op, errMsg, pages.SignInBase, pages.SignIn)
}

func renderError(ctx context.Context, w http.ResponseWriter, logger *zap.SugaredLogger, op, errMsg string, basePage func(templ.Component) templ.Component, authPage func(string) templ.Component) {
	if err := render.Render(ctx, w, basePage(authPage(errMsg))); err != nil {
		logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
