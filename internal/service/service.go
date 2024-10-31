package service

import (
	"constructor/internal/model"
	"context"
	"errors"

	"github.com/go-playground/validator"
)

var ErrUndefinedRole = errors.New("undefined role")

type authStorager interface {
	SignUp(ctx context.Context, user model.SignUpUser) error
	SignIn(ctx context.Context, user model.SignInUser) (int, string, error)
}

type courseStorager interface {
	CheckAndExpireCourseAccess() error
}

type authService struct {
	authStore  authStorager
	validator  *validator.Validate
	passSecret string
	jwtSecret  string
}

type courseService struct {
	courseStore courseStorager
}

func NewAuthService(passSecret, jwtSecret string, authStorage authStorager) *authService {
	return &authService{
		authStore:  authStorage,
		validator:  validator.New(),
		passSecret: passSecret,
		jwtSecret:  jwtSecret,
	}
}

func NewCourseService(courseStorage courseStorager) *courseService {
	return &courseService{
		courseStore: courseStorage,
	}
}
