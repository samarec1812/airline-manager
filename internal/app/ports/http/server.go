package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/samarec1812/airline-manager/internal/app/service"
	"github.com/samarec1812/airline-manager/internal/config"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

func NewHTTPServer(cfg config.HTTPServer, logger *slog.Logger, a service.App) *http.Server {
	handler := chi.NewRouter()

	s := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout * time.Second,
		WriteTimeout: cfg.Timeout * time.Second,
		IdleTimeout:  cfg.IdleTimeout * time.Second,
	}

	AppRouter(handler, logger, a)

	return s
}