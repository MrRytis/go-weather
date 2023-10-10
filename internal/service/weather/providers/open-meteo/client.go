package openMeteo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type coordination struct {
	Longitude string
	Latitude  string
}

type currentWeather struct {
	Time        string  `json:"time"`
	Temp        float32 `json:"temperature"`
	WindSpeed   float32 `json:"windspeed"`
	WindDir     int32   `json:"winddirection"`
	WeatherCode int32   `json:"weathercode"`
}

type daily struct {
	PrecipitationSum []float32 `json:"precipitation_sum"`
}

type response struct {
	Weather currentWeather `json:"current_weather"`
	Daily   daily          `json:"daily"`
}

func getWeather(city string) (*response, error) {
	coord := convertCityToStation(city)

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m,dewpoint_2m,precipitation_probability,rain&daily=weathercode,precipitation_sum&current_weather=true&timezone=GMT&forecast_days=1",
		coord.Latitude,
		coord.Longitude)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}
	defer res.Body.Close()

	var meteoRes response
	err = json.NewDecoder(res.Body).Decode(&meteoRes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode httpUtils body: %w", err)
	}

	return &meteoRes, nil
}

func convertCityToStation(city string) coordination {
	switch city {
	case "VILNIUS":
		return coordination{
			Longitude: "25.2798",
			Latitude:  "54.6892",
		}
	case "KAUNAS":
		return coordination{
			Longitude: "23.9096",
			Latitude:  "54.9027",
		}
	case "KLAIPEDA":
		return coordination{
			Longitude: "21.1391",
			Latitude:  "55.7068",
		}
	default:
		return coordination{}
	}
}
