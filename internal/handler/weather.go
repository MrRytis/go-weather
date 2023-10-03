package handler

import (
	"github.com/MrRytis/go-weather/internal/model"
	weatherService "github.com/MrRytis/go-weather/internal/service/weather"
	"github.com/MrRytis/go-weather/pkg/httpUtils"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func (h *Handler) GetCitiesWeatherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := strings.ToUpper(vars["city"])

	if !weatherService.IsCitySupported(city) {
		httpUtils.ErrorJSON(w, "City is not supported", http.StatusBadRequest)
	}

	weather, err := weatherService.GetWeather(city)
	if err != nil {
		httpUtils.ErrorJSON(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var res []model.WeatherResponse
	for _, v := range *weather {
		res = append(res, model.WeatherResponse{
			Temp:          v.Temp,
			FeelsLike:     v.FeelsLike,
			WindSpeed:     v.WindSpeed,
			WindDeg:       v.WindDeg,
			Clouds:        v.Clouds,
			Pressure:      v.Pressure,
			Humidity:      v.Humidity,
			Precipitation: v.Precipitation,
			Weather:       v.Weather,
			Provider:      v.Provider,
			City:          v.City,
		})
	}

	httpUtils.JSON(w, http.StatusOK, res)
}

func (h *Handler) GetHistoricalCitiesWeatherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	page := httpUtils.GetIntQueryParam(r, "page", 1) - 1
	limit := httpUtils.GetIntQueryParam(r, "limit", 10)

	if !weatherService.IsCitySupported(city) {
		httpUtils.ErrorJSON(w, "City is not supported", http.StatusBadRequest)
	}

	weather, err := h.Repository.GetHistoricalWeatherPaginate(city, page, limit)
	if err != nil {
		httpUtils.ErrorJSON(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	total, err := h.Repository.GetHistoricalWeatherCount(city)
	if err != nil {
		httpUtils.ErrorJSON(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var data []model.WeatherResponse
	for _, v := range weather {
		data = append(data, model.WeatherResponse{
			Temp:          v.Temp,
			FeelsLike:     v.FeelsLike,
			WindSpeed:     v.WindSpeed,
			WindDeg:       v.WindDeg,
			Clouds:        v.Clouds,
			Pressure:      v.Pressure,
			Humidity:      v.Humidity,
			Precipitation: v.Precipitation,
			Weather:       v.Weather,
			Provider:      v.Provider,
		})
	}

	res := model.HistoricalWeatherResponse{
		Data:  data,
		Page:  page,
		Total: total,
		Limit: limit,
	}

	httpUtils.JSON(w, http.StatusOK, res)
}
