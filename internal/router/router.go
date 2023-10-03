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
	h := handler.NewHandler(app.Db, app.C)

	r.HandleFunc("/", h.IndexHandler).Methods(http.MethodGet)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(m.RateLimitMiddleware)

	// Auth
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)
	auth.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)

	// weather
	weather := api.PathPrefix("/weather").Subrouter()
	weather.Use(m.AuthMiddleware)
	weather.HandleFunc("/{city}/now", h.GetCitiesWeatherHandler).Methods(http.MethodGet)
	weather.HandleFunc("/{city}/history", h.GetHistoricalCitiesWeatherHandler).Methods(http.MethodGet)

	// cron
	cron := api.PathPrefix("/cron").Subrouter()
	cron.Use(m.AuthMiddleware)
	cron.HandleFunc("/start", h.StartCronHandler).Methods(http.MethodPost)
	cron.HandleFunc("/stop", h.StopCronHandler).Methods(http.MethodPost)
	cron.HandleFunc("/add", h.AddCronJobHandler).Methods(http.MethodPost)

	return r
}
