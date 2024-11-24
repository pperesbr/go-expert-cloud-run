package application

import (
	"github.com/pperesbr/go-expert-cloud-run/internal/application/ports"
	"github.com/pperesbr/go-expert-cloud-run/internal/domain"
)

type WeatherService struct {
	zipService ports.ZipService
	weatherAPI ports.WeatherAPI
}

func NewWeatherService(zipService ports.ZipService, weatherAPI ports.WeatherAPI) *WeatherService {
	return &WeatherService{
		zipService: zipService,
		weatherAPI: weatherAPI,
	}
}

func (ws *WeatherService) GetWeather(request InputRequestDTO) (*WeatherResponseDTO, error) {

	zipInfo, err := ws.zipService.SearchZipInfo(request.ZipCode)
	if err != nil {
		return nil, err
	}

	celsius, err := ws.weatherAPI.GetTemperatureByCityAndState(zipInfo.City, zipInfo.State)
	if err != nil {
		return nil, err
	}

	temperature := domain.NewTemperature(celsius)

	return &WeatherResponseDTO{
		Celsius:    temperature.Celsius,
		Fahrenheit: temperature.Fahrenheit,
		Kelvin:     temperature.Kelvin,
	}, nil
}
