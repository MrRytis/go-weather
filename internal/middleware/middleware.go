package middleware

import (
	"github.com/MrRytis/go-weather/internal/storage"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type Middleware struct {
	Repository *storage.Repository
	Limiter    *rate.Limiter
}

func NewMiddleware(db *gorm.DB) *Middleware {
	return &Middleware{
		Repository: storage.NewRepository(db),
		Limiter:    rate.NewLimiter(0.2, 1),
	}
}
