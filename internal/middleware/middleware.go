package middleware

import (
	"github.com/MrRytis/go-weather/internal/storage"
	"gorm.io/gorm"
)

type Middleware struct {
	Repository *storage.Repository
}

func NewMiddleware(db *gorm.DB) *Middleware {
	return &Middleware{
		Repository: storage.NewRepository(db),
	}
}
