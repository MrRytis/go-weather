package jobs

import (
	weatherService "github.com/MrRytis/go-weather/internal/service/weather"
	"gorm.io/gorm"
	"log"
)

type WeatherCronJob struct{}

func (j *WeatherCronJob) GetName() string {
	return "weather"
}

func (j *WeatherCronJob) RunJob(db *gorm.DB) {
	log.Printf("running cron job %s", j.GetName())

	cities := weatherService.GetAllowedCities()

	for _, city := range cities {
		weather, err := weatherService.GetWeather(city)
		if err != nil {
			log.Printf("failed to get weather data for city %s: %s", city, err.Error())
		}

		for _, w := range *weather {
			db.Create(&w)
		}
	}

	log.Printf("cron job %s finished", j.GetName())
}
