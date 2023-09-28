package model

type HistoricalWeatherResponse struct {
	Data  []WeatherResponse `json:"data"`
	Page  int               `json:"page"`
	Total int64             `json:"total"`
	Limit int               `json:"limit"`
}

type WeatherResponse struct {
	City          string  `json:"city"`
	Temp          float32 `json:"temp"`
	FeelsLike     float32 `json:"feels_like"`
	Pressure      float32 `json:"pressure"`
	Humidity      int32   `json:"humidity"`
	WindSpeed     float32 `json:"wind_speed"`
	WindDeg       int32   `json:"wind_deg"`
	Clouds        int32   `json:"clouds"`
	Weather       string  `json:"weather"`
	Precipitation int32   `json:"precipitation"`
	Provider      string  `json:"provider"`
}
