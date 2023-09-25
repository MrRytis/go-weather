package middleware

import (
	"github.com/MrRytis/go-weather/internal/service"
	"github.com/MrRytis/go-weather/pkg/response"
	"net/http"
	"strings"
)

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		if tokenString == "" {
			response.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := service.ParseJWT(tokenString)
		if err != nil {
			response.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// can set user id to context here

		next.ServeHTTP(w, r)
	})
}
