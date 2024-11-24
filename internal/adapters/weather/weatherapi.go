package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pperesbr/go-expert-cloud-run/internal/application/ports"
	"github.com/pperesbr/go-expert-cloud-run/pkg"
	"io"
	"net/http"
)

var ErrWeatherAPIKey = errors.New("weather api key is invalid or missing")

type WeatherAPI struct {
	APIKey string
}

func NewWeatherAPI(apiKey string) *WeatherAPI {
	return &WeatherAPI{APIKey: apiKey}
}

func (w *WeatherAPI) GetTemperatureByCityAndState(city string, state string) (float64, error) {
	url := fmt.Sprintf(
		"https://api.weatherapi.com/v1/current.json?key=%s&q=%s,%s&aqi=no",
		w.APIKey,
		pkg.RemoveAccents(city),
		pkg.RemoveAccents(state),
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			return 0, ErrWeatherAPIKey
		}

		return 0, ports.ErrWeatherNotFound
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data["current"].(map[string]interface{})["temp_c"].(float64), nil
}

var _ ports.WeatherAPI = &WeatherAPI{}
