package main

import (
	"constructor/internal/config"
	httpserver "constructor/internal/http-server"
	"constructor/internal/lib/logger"
	"constructor/internal/service"
	"constructor/internal/storage/postgres"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	var (
		logger = must(logger.New())
		config = must(config.New(".env"))

		conn      = must(postgres.New(config.PGConfig.DSN))
		authDB    = postgres.NewAuthStorage(conn)
		courseyDB = postgres.NewCourseStorage(conn)

		authService   = service.NewAuthService(config.PASS_SECRET, config.JWT_SECRET, authDB)
		courseService = service.NewCourseService(courseyDB)

		r = chi.NewRouter()
	)
	defer conn.Close()

	// Initializes middlewares for all server requests,
	// other middlewares can be initialized in the New function, see httpserver.New().
	r.Use(
		middleware.RequestID,
		middleware.Logger,
		httprate.Limit(5, time.Second),
	)

	// Initializes server routes and returns a completed http server.
	server := httpserver.New(r, logger, config, authService)

	courseService.StartCourseAccessExpirationCheck()

	logger.Infof("runnig server at: %s", config.HTTPConfig.Addr)
	if err := server.ListenAndServe(); err != nil {
		logger.Errorf("server shutdown with error: %w", err)
		os.Exit(1)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
