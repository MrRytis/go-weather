package meteo

import (
	"fmt"
	"github.com/MrRytis/go-weather/internal/storage"
)

const ProviderName = "meteo"

type Provider struct{}

func (p Provider) GetCitiesWeather(city string) (*[]storage.Weather, error) {
	res, err := getWeather(city)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}

	var w []storage.Weather
	for _, v := range *res {
		w = append(w, storage.Weather{
			Time:          v.Time,
			Temp:          v.Temp,
			FeelsLike:     v.FeelsLike,
			WindSpeed:     v.WindSpeed,
			WindDeg:       v.WindDir,
			Clouds:        v.Clouds,
			Pressure:      v.Pressure,
			Humidity:      v.Humidity,
			Precipitation: v.Precipitation,
			Weather:       v.Condition,
			Provider:      ProviderName,
		})
	}

	return &w, nil
}

func (p Provider) GetCurrentWeather(city string) (*storage.Weather, error) {
	res, err := getWeather(city)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}

	v := (*res)[len(*res)-1] //latest weather data

	w := storage.Weather{
		Time:          v.Time,
		Temp:          v.Temp,
		FeelsLike:     v.FeelsLike,
		WindSpeed:     v.WindSpeed,
		WindDeg:       v.WindDir,
		Clouds:        v.Clouds,
		Pressure:      v.Pressure,
		Humidity:      v.Humidity,
		Precipitation: v.Precipitation,
		Weather:       v.Condition,
	}

	return &w, nil
}

func (p Provider) GetName() string {
	return ProviderName
}
