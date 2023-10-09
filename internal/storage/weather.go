package storage

import (
	"time"
)

type Weather struct {
	Id            int32     `gorm:"primaryKey"`
	City          string    `gorm:"not null"`
	Temp          float32   `gorm:"not null"`
	FeelsLike     float32   `gorm:"not null"`
	WindSpeed     float32   `gorm:"not null"`
	WindDeg       int32     `gorm:"not null"`
	Weather       string    `gorm:"not null"`
	Precipitation int32     `gorm:"not null"`
	Time          time.Time `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	Provider      string    `gorm:"not null"`
	Pressure      *float32
	Humidity      *int32
	Clouds        *int32
}

func (Weather) TableName() string {
	return "weather_data"
}

func (r *Repository) CreateWeather(w Weather) error {
	return r.Db.Create(&w).Error
}

func (r *Repository) GetWeatherByCity(city string) ([]Weather, error) {
	var w []Weather
	err := r.Db.Where("city = ?", city).Find(&w).Error

	return w, err
}

func (r *Repository) GetHistoricalWeatherPaginate(city string, page int, limit int) ([]Weather, error) {
	var w []Weather
	err := r.Db.Where("city = ?", city).Offset(page * limit).Limit(limit).Find(&w).Error

	return w, err
}

func (r *Repository) GetHistoricalWeatherCount(city string) (int64, error) {
	var count int64
	err := r.Db.Model(&Weather{}).Where("city = ?", city).Count(&count).Error

	return count, err
}
