package handler

import (
	"github.com/MrRytis/go-weather/internal/storage"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Handler struct {
	Repository *storage.Repository
	Cron       *cron.Cron
}

func NewHandler(db *gorm.DB, c *cron.Cron) *Handler {
	return &Handler{
		Repository: storage.NewRepository(db),
		Cron:       c,
	}
}
