package internal

import (
	"github.com/MrRytis/go-weather/internal/storage"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type App struct {
	Db *gorm.DB
	C  *cron.Cron
}

func NewApp() *App {
	db := storage.NewDb()
	c := cron.New()

	return &App{
		Db: db,
		C:  c,
	}
}
