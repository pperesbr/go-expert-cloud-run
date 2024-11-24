package main

import (
	"fmt"
	"github.com/pperesbr/go-expert-cloud-run/config"
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/api"
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/weather"
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/zip"
	"github.com/pperesbr/go-expert-cloud-run/internal/application"
	"log"
	"net/http"
	"os"
)

func main() {

	appConfig := config.NewAppConfig(
		os.Getenv("SERVER_PORT"),
		os.Getenv("WEATHER_API_KEY"),
	)

	service := application.NewWeatherService(zip.NewViaCep(), weather.NewWeatherAPI(appConfig.WeatherAPIKey))
	api := api.NewApi(service)
	http.HandleFunc("/{zipcode}", api.GetWeather)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", appConfig.ServerPort), nil))
}
