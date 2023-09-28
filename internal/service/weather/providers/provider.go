package weatherProviders

import (
	"github.com/MrRytis/go-weather/internal/service/weather/providers/meteo"
	"github.com/MrRytis/go-weather/internal/storage"
)

type Provider interface {
	GetCitiesWeather(city string) (*[]storage.Weather, error)
	GetCurrentWeather(city string) (*storage.Weather, error)
	GetName() string
}

func GetProviders() []Provider {
	return []Provider{
		&meteo.Provider{},
	}
}
