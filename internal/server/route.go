package server

import (
	"net/http"

	"github.com/heismyke/local_business_booking_app/internal/handler"
	"github.com/heismyke/local_business_booking_app/internal/service"
)


func (s *Server) RegisterRoutes() http.Handler{
  mux := http.NewServeMux()
  service := service.New()
  handler := handler.New(service)

  mux.Handle("/api/v1/street_business/auth_v1", http.HandlerFunc(handler.CreateUser))
  return s.corsMiddleware(mux)
}



func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
