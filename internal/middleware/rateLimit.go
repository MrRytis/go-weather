package middleware

import (
	"github.com/MrRytis/go-weather/pkg/httpUtils"
	"net/http"
)

func (m *Middleware) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.Limiter.Allow() == false {
			httpUtils.ErrorJSON(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
