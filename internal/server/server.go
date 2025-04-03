package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server структура HTTP сервера
type Server struct {
	logger *log.Logger
	srv    *http.Server
}

// New создает и настраивает новый экземпляр сервера
func NewServer(logger *log.Logger) *Server {
	router := createRouter()

	return &Server{
		logger: logger,
		srv: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

// createRouter создает и настраивает роутер
func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.HandlerRoot)
	router.HandleFunc("POST /upload", handlers.UploadHandler)

	return router
}

// Start запускает HTTP сервер
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}
