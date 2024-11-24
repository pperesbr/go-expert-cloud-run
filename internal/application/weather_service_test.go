package application

import (
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/weather"
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/zip"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"testing"
)

func TestWeatherService_GetWeather(t *testing.T) {

	assert := assert.New(t)

	service := NewWeatherService(zip.NewViaCep(), weather.NewWeatherAPI(os.Getenv("WEATHER_API_KEY")))
	temperature, err := service.GetWeather(InputRequestDTO{ZipCode: "87560000"})
	assert.NoError(err)

	scale := math.Pow(10, float64(1))
	assert.Equal(temperature.Celsius, 28.6)
	assert.Equal(temperature.Fahrenheit, math.Round(temperature.Celsius*1.8+32*scale)/scale)
	assert.Equal(temperature.Kelvin, math.Round(temperature.Celsius+273.15*scale)/scale)
}
