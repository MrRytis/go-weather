package middleware

import (
	"github.com/MrRytis/go-weather/internal/service/auth"
	"github.com/MrRytis/go-weather/pkg/httpUtils"
	"net/http"
	"strings"
)

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		if tokenString == "" {
			httpUtils.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := authService.ParseJWT(tokenString)
		if err != nil {
			httpUtils.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// can set user id to context here

		next.ServeHTTP(w, r)
	})
}
