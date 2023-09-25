package router

import (
	"github.com/MrRytis/go-weather/internal"
	"github.com/MrRytis/go-weather/internal/handler"
	"github.com/MrRytis/go-weather/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(app *internal.App) *mux.Router {
	r := mux.NewRouter()

	m := middleware.NewMiddleware(app.Db)
	h := handler.NewHandler(app.Db)

	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)

	// weather
	weather := api.PathPrefix("/weather").Subrouter()
	weather.Use(m.AuthMiddleware)
	weather.HandleFunc("/:city/today", h.GetCitiesWeatherHandler).Methods(http.MethodGet)
	weather.HandleFunc("/:city/history", h.GetHistoricalCitiesWeatherHandler).Methods(http.MethodGet)

	return r
}
