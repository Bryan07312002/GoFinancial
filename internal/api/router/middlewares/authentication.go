package middlewares

import (
	"financial/internal/services"
	"financial/internal/sessions"

	"github.com/gorilla/mux"

	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserKey contextKey = "UserID"

type IsAuthenticatedFactory interface {
	CreateIsAuthenticated() services.IsAuthenticated
}

func CreateAuthMiddleware(factory IsAuthenticatedFactory) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authInput := strings.Split(r.Header.Get("Authorization"), " ")
			if len(authInput) <= 1 {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := authInput[1]

			userID, isAuthenticated := factory.CreateIsAuthenticated().
				Run(sessions.Token(token))
			if !isAuthenticated {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, userID)
			updatedReq := r.WithContext(ctx)

			next.ServeHTTP(w, updatedReq)
		})
	}
}
