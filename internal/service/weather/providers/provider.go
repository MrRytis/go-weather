package weatherProviders

import (
	"github.com/MrRytis/go-weather/internal/service/weather/providers/meteo"
	openMeteo "github.com/MrRytis/go-weather/internal/service/weather/providers/open-meteo"
	"github.com/MrRytis/go-weather/internal/storage"
)

type Provider interface {
	GetCurrentWeather(city string) (*storage.Weather, error)
	GetName() string
}

func GetProviders() []Provider {
	return []Provider{
		&meteo.Provider{},
		&openMeteo.Provider{},
	}
}
