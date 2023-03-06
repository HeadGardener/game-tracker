package server

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     zap.Logger
}

func New(logger zap.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
