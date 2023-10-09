package openMeteo

import (
	"fmt"
	"github.com/MrRytis/go-weather/internal/storage"
	"time"
)

const ProviderName = "open-meteo"

type Provider struct{}

func (p Provider) GetCurrentWeather(city string) (*storage.Weather, error) {
	res, err := getWeather(city)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}

	if res.Weather.Time == "" {
		res.Weather.Time = time.Now().Format("2006-01-02T15:04")

	}

	t, err := time.Parse("2006-01-02T15:04", res.Weather.Time)
	if err != nil {
		return nil, fmt.Errorf("failed to parse time: %w", err)
	}

	participation := int32(0)
	if len(res.Daily.PrecipitationSum) > 0 {
		participation = int32(res.Daily.PrecipitationSum[0])
	}

	w := storage.Weather{
		Time:          t,
		Temp:          res.Weather.Temp,
		FeelsLike:     res.Weather.Temp,
		WindSpeed:     res.Weather.WindSpeed,
		WindDeg:       res.Weather.WindDir,
		Clouds:        nil,
		Pressure:      nil,
		Humidity:      nil,
		Precipitation: participation,
		Weather:       translateVMOCodeToText(res.Weather.WeatherCode),
		Provider:      ProviderName,
		City:          city,
	}

	return &w, nil
}

func (p Provider) GetName() string {
	return ProviderName
}

func translateVMOCodeToText(code int32) string {
	codes := map[int32]string{
		00: "No significant weather observed",
		01: "Clouds generally dissolving or becoming less developed during the past hour",
		02: "State of sky on the whole unchanged during the past hour",
		03: "Clouds generally forming or developing during the past hour",
		04: "Haze or smoke, or dust in suspension in the air, visibility equal to, or greater than, 1 km",
		05: "Haze or smoke, or dust in suspension in the air, visibility less than 1 km",
		10: "Mist",
		11: "Diamond dust",
		12: "Distant lightning",
		18: "Squalls",
		20: "Fog",
		21: "PRECIPITATION",
		22: "Drizzle (not freezing) or snow grains",
		23: "Rain (not freezing)",
		24: "Snow",
		25: "Freezing drizzle or freezing rain",
		26: "Thunderstorm (with or without precipitation)",
		27: "BLOWING OR DRIFTING SNOW OR SAND",
		28: "Blowing or drifting snow or sand, visibility equal to, or greater than, 1 km",
		29: "Blowing or drifting snow or sand, visibility less than 1 km",
		30: "FOG",
		31: "Fog or ice fog in patches",
		32: "Fog or ice fog, has become thinner during the past hour",
		33: "Fog or ice fog, no appreciable change during the past hour",
		34: "Fog or ice fog, has begun or become thicker during the past hour",
		40: "PRECIPITATION",
		41: "Precipitation, slight or moderate",
		42: "Precipitation, heavy",
		43: "Liquid precipitation, slight or moderate",
		44: "Liquid precipitation, heavy",
		45: "Solid precipitation, slight or moderate",
		46: "Solid precipitation, heavy",
		47: "Freezing precipitation, slight or moderate",
		48: "Freezing precipitation, heavy",
		50: "DRIZZLE",
		51: "Drizzle, not freezing, slight",
		52: "Drizzle, not freezing, moderate",
		53: "Drizzle, not freezing, heavy",
		54: "Drizzle, freezing, slight",
		55: "Drizzle, freezing, moderate",
		56: "Drizzle, freezing, heavy",
		57: "Drizzle and rain, slight",
		58: "Drizzle and rain, moderate or heavy",
		60: "RAIN",
		61: "Rain, not freezing, slight",
		62: "Rain, not freezing, moderate",
		63: "Rain, not freezing, heavy",
		64: "Rain, freezing, slight",
		65: "Rain, freezing, moderate",
		66: "Rain, freezing, heavy",
		67: "Rain (or drizzle) and snow, slight",
		68: "Rain (or drizzle) and snow, moderate or heavy",
		70: "SNOW",
		71: "Snow, slight",
		72: "Snow, moderate",
		73: "Snow, heavy",
		74: "Ice pellets, slight",
		75: "Ice pellets, moderate",
		76: "Ice pellets, heavy",
		77: "Snow grains",
		78: "Ice crystals",
		80: "SHOWER(S) or INTERMITTENT PRECIPITATION",
		81: "Rain shower(s) or intermittent rain, slight",
		82: "Rain shower(s) or intermittent rain, moderate",
		83: "Rain shower(s) or intermittent rain, heavy",
		84: "Rain shower(s) or intermittent rain, violent",
		85: "Snow shower(s) or intermittent snow, slight",
		86: "Snow shower(s) or intermittent snow, moderate",
		87: "Snow shower(s) or intermittent snow, heavy",
		89: "Hail",
		90: "THUNDERSTORM",
		91: "Thunderstorm, slight or moderate, with no precipitation",
		92: "Thunderstorm, slight or moderate, with rain showers and/or snow showers",
		93: "Thunderstorm, slight or moderate, with hail",
		94: "Thunderstorm, heavy, with no precipitation",
		95: "Thunderstorm, heavy, with rain showers and/or snow showers",
		96: "Thunderstorm, heavy, with hail",
		99: "Tornado",
	}

	if description, exists := codes[code]; exists {
		return description
	}

	return "-"
}
