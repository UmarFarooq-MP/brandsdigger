package middleware

import (
	"context"
	"net/http"
	"strings"

	"brandsdigger/internal/domain/auth" // Use the interface, not the concrete implementation
)

type contextKey string

const ClaimsContextKey = contextKey("claims")

// JWTMiddleware returns a middleware that validates JWT tokens using the provided TokenService
func JWTMiddleware(tokenService auth.TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := tokenService.ValidateToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
				return
			}

			// Add claims to request context
			ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
