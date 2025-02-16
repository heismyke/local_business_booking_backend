package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/heismyke/local_business_booking_app/internal/service"
)

// Config holds configuration settings for the server.
type Config struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Server holds the dependencies for the HTTP server.
type Server struct {
	port int
	db     service.Service
}

// NewServer creates a new HTTP server with the given configuration and dependencies.
func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	s := &Server{
    port : port,
		db:     service.New(),
	}

	return &http.Server{
    Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,

	}
}

