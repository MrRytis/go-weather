package jobs

import "gorm.io/gorm"

type WeatherCronJob struct{}

func (j *WeatherCronJob) GetName() string {
	return "weather"
}

func (j *WeatherCronJob) RunJob(db *gorm.DB) {
	// TODO: implement
}
