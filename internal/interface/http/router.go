package http

import (
	"brandsdigger/internal/domain/auth"
	"brandsdigger/internal/interface/http/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateRouter sets up the entire router with public and protected routes
func CreateRouter(authHandler *AuthHandler, namesHandler *NamesHandler, tokenService auth.TokenService) *mux.Router {
	r := mux.NewRouter()

	// ✅ Public routes
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/signup", authHandler.Signup).Methods("POST")

	// ✅ Protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTMiddleware(tokenService))

	// Add JWT-protected endpoints here
	protected.HandleFunc("/generate/names", namesHandler.ServeHTTP).Methods("POST")

	return r
}
