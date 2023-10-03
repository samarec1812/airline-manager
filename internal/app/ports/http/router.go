package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
account
	"github.com/samarec1812/airline-manager/internal/app/service"
)

func AppRouter(r *chi.Mux, logger *slog.Logger, a service.App) {
	r.Use(middleware.RequestID)
	r.Use(LoggerMiddleware(logger))
	r.Use(middleware.Recoverer)

	r.Post("/airline", CreateAirline(logger, a))
	r.Delete("/airline", DeleteAirline(logger, a))

}
