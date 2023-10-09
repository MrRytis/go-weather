package meteo

import (
	"fmt"
	"github.com/MrRytis/go-weather/internal/storage"
	"time"
)

const ProviderName = "meteo"

type Provider struct{}

func (p Provider) GetCurrentWeather(city string) (*storage.Weather, error) {
	res, err := getWeather(city)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}

	v := (*res)[len(*res)-1] //latest weather data

	t, err := time.Parse("2006-01-02 15:04:05", v.Time)
	if err != nil {
		return nil, fmt.Errorf("failed to parse time: %w", err)
	}

	w := storage.Weather{
		Time:          t,
		Temp:          v.Temp,
		FeelsLike:     v.FeelsLike,
		WindSpeed:     v.WindSpeed,
		WindDeg:       v.WindDir,
		Clouds:        &v.Clouds,
		Pressure:      &v.Pressure,
		Humidity:      &v.Humidity,
		Precipitation: int32(v.Precipitation),
		Weather:       v.Condition,
		Provider:      ProviderName,
		City:          city,
	}

	return &w, nil
}

func (p Provider) GetName() string {
	return ProviderName
}
