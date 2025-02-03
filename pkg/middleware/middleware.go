package middleware

import (
	"context"
	"golang_starter_template/pkg/session"
	"golang_starter_template/pkg/utils"
	"log"
	"net/http"
	"strings"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request from", r.RemoteAddr, "to", r.URL)
		utils.LoggerInfo.Println(utils.Info+"Request from", r.RemoteAddr, "to", r.URL, utils.Reset)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.TrimSpace(authHeader) == "" {
			utils.LoggerError.Println(utils.Error+"Authorization header required", utils.Reset)
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			utils.LoggerError.Println(utils.Error+"Invalid token format", utils.Reset)
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		userID, roles, err := session.ValidateToken(tokenParts[1])
		if err != nil {
			utils.LoggerError.Println(utils.Error+"Invalid token", utils.Reset)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Set user id and roles to request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "roles", roles)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
