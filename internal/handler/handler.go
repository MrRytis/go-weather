package handler

import (
	"github.com/MrRytis/go-weather/internal/storage"
	"gorm.io/gorm"
)

type Handler struct {
	Repository *storage.Repository
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		Repository: storage.NewRepository(db),
	}
}
