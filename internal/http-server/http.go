package httpserver

import (
	"constructor/internal/config"
	"constructor/internal/http-server/handlers/auth"
	"constructor/internal/http-server/handlers/constructor"
	"constructor/internal/http-server/handlers/course"
	"constructor/internal/http-server/handlers/home"
	"constructor/internal/http-server/middlewares"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

const readHeaderTimeout = 10 * time.Second

func initHandlers(r *chi.Mux, logger *zap.SugaredLogger, authService auth.AuthService, courseService course.CourseService) {
	r.Get("/", home.HomePage(logger))

	r.Get("/signin", auth.SignInPage(logger, authService))
	r.Post("/signin", auth.SignInPage(logger, authService))

	r.Get("/signup", auth.SignUpPage(logger, authService))
	r.Post("/signup", auth.SignUpPage(logger, authService))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/constructor", constructor.ConstructorPage(logger, courseService))
		r.Post("/constructor", constructor.ConstructorPage(logger, courseService))

		r.Get("/course/create", course.CourseCreationPage(logger, courseService))
		r.Post("/course/create", course.CourseCreationPage(logger, courseService))

		r.Get("/course/{id}", course.CoursePage(logger, courseService))
		r.Post("/course/{id}", course.CoursePage(logger, courseService))

		// code below needs to refactor
		// r.Post("/courses", course.CoursesList())
	})
}

func New(r *chi.Mux, logger *zap.SugaredLogger, config *config.Config, authService auth.AuthService, courseService course.CourseService) *http.Server {

	initHandlers(r, logger, authService, courseService)

	httpSrv := &http.Server{
		Addr:              config.HTTPConfig.Addr,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return httpSrv
}
