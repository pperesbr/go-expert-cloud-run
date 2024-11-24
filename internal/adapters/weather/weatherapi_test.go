package weather

import (
	"github.com/pperesbr/go-expert-cloud-run/internal/application/ports"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWeatherAPI_GetTemperatureByCityAndState(t *testing.T) {

	assert := assert.New(t)
	weatherAPI := NewWeatherAPI("fake-api-key")

	_, err := weatherAPI.GetTemperatureByCityAndState("Ipora", "PR")
	assert.ErrorIs(err, ErrWeatherAPIKey)

	weatherAPI.APIKey = os.Getenv("WEATHER_KEY")
	_, err = weatherAPI.GetTemperatureByCityAndState("invalid", "")
	assert.ErrorIs(err, ports.ErrWeatherNotFound)

	temperature, err := weatherAPI.GetTemperatureByCityAndState("Ipora", "PR")
	assert.NoError(err)
	assert.Equal(temperature, 23.4)

}
