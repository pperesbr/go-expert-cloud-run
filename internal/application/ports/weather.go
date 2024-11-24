package ports

import "errors"

var ErrWeatherNotFound = errors.New("weather not found")

type WeatherAPI interface {
	GetTemperatureByCityAndState(city string, state string) (float64, error)
}
