package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Weather []weather `json:"observations"`
}

type weather struct {
	Time          string  `json:"observationTimeUtc"`
	Temp          float32 `json:"airTemperature"`
	FeelsLike     float32 `json:"feelsLikeTemperature"`
	WindSpeed     float32 `json:"windSpeed"`
	WindGust      float32 `json:"windGust"`
	WindDir       int32   `json:"windDirection"`
	Clouds        int32   `json:"cloudCover"`
	Pressure      float32 `json:"seaLevelPressure"`
	Humidity      int32   `json:"relativeHumidity"`
	Precipitation float32 `json:"precipitation"`
	Condition     string  `json:"conditionCode"`
}

func getWeather(city string) (*[]weather, error) {
	url := fmt.Sprintf("https://api.meteo.lt/v1/stations/%s/observations/latest", convertCityToStation(city))

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

	return &meteoRes.Weather, nil
}

func convertCityToStation(city string) string {
	switch city {
	case "VILNIUS":
		return "vilniaus-ams"
	case "KAUNAS":
		return "kauno-ams"
	case "KLAIPEDA":
		return "klaipedos-ams"
	default:
		return ""
	}
}
