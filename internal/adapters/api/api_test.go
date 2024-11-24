package api

import (
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/weather"
	"github.com/pperesbr/go-expert-cloud-run/internal/adapters/zip"
	"github.com/pperesbr/go-expert-cloud-run/internal/application"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_GetWeather(t *testing.T) {

	assert := assert.New(t)
	service := application.NewWeatherService(zip.NewViaCep(), weather.NewWeatherAPI(os.Getenv("WEATHER_API_KEY")))
	api := NewApi(service)
	mux := http.NewServeMux()
	mux.Handle("/{zipcode}", http.HandlerFunc(api.GetWeather))

	req := httptest.NewRequest("GET", "/87560-000", nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)
	response := recorder.Result()
	assert.Equal(422, response.StatusCode)

	req = httptest.NewRequest("GET", "/12345678", nil)
	recorder = httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)
	response = recorder.Result()
	assert.Equal(404, response.StatusCode)

	req = httptest.NewRequest("GET", "/87560000", nil)
	recorder = httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)
	response = recorder.Result()
	assert.Equal(200, response.StatusCode)
}
