package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Garik-/ws-epoll/internal/zlog"
)

const (
	shutdownTimeout   = 5 * time.Second
	readTimeout       = 1 * time.Second
	writeTimeout      = 1 * time.Second
	idleTimeout       = 30 * time.Second
	readHeaderTimeout = 2 * time.Second
)

type Service struct {
	server *http.Server
}

func New(config *Config) *Service {
	mux := http.NewServeMux()
	mux.HandleFunc("/", wsUpgrade)

	return &Service{
		server: &http.Server{
			Addr:              config.Addr,
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
			IdleTimeout:       idleTimeout,
			ReadHeaderTimeout: readHeaderTimeout,
			Handler:           mux,
		},
	}
}

func (s *Service) Run(_ context.Context) error {
	zlog.Info("ListenAndServe", zlog.String("addr", s.server.Addr))

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server error: %w", err)
	}

	return nil
}

func (s *Service) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("service close error: %w", err)
	}

	return nil
}
