package internal

import (
	"github.com/MrRytis/go-weather/internal/storage"
	"gorm.io/gorm"
)

type App struct {
	Db *gorm.DB
}

func NewApp() *App {
	db := storage.NewDb()

	return &App{
		Db: db,
	}
}
