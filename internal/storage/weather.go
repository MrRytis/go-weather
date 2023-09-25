package storage

import "time"

type Weather struct {
	Id        int32
	City      string
	Temp      float32
	FeelsLike float32
	MinTemp   float32
	MaxTemp   float32
	Pressure  int32
	Humidity  int32
	WindSpeed float32
	WindDeg   int32
	Clouds    int32
	Weather   string
	CreateAt  time.Time
}
