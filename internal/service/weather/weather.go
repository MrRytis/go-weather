package weatherService

import (
	"fmt"
	weatherProviders "github.com/MrRytis/go-weather/internal/service/weather/providers"
	"github.com/MrRytis/go-weather/internal/storage"
	"log"
	"sync"
)

var supportedCities = []string{"VILNIUS", "KAUNAS", "KLAIPEDA"}

func GetWeather(city string) (*[]storage.Weather, error) {
	if !IsCitySupported(city) {
		return nil, fmt.Errorf("city %s is not supported", city)
	}

	providers := weatherProviders.GetProviders()

	resultChan := make(chan *storage.Weather, len(providers))
	var wg sync.WaitGroup

	// Launch goroutines
	for _, p := range weatherProviders.GetProviders() {
		wg.Add(1) // Increment the WaitGroup counter
		go func(p weatherProviders.Provider, resultChan chan *storage.Weather) {
			defer wg.Done() // Decrement the WaitGroup counter when done
			w, err := p.GetCurrentWeather(city)
			if err != nil {
				log.Printf("failed to get weather data from provider %s: %s", p.GetName(), err.Error())
			}

			resultChan <- w
		}(p, resultChan)
	}

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(resultChan)
	}()

	// Collect results from the channel
	var res []storage.Weather
	for result := range resultChan {
		res = append(res, *result)
	}

	return &res, nil
}

func IsCitySupported(city string) bool {
	for _, c := range supportedCities {
		if c == city {
			return true
		}
	}

	return false
}
