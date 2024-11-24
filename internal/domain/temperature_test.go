package domain

import (
	"math"
	"testing"
)

func TestNewTemperature(t *testing.T) {

	scale := math.Pow(10, float64(1))
	celsius := 25.0
	temperature := NewTemperature(celsius)

	expectedFahrenheit := math.Round(25*1.8+32*scale) / scale
	expectedKelvin := math.Round(celsius+273.15*scale) / scale

	if got := temperature.Fahrenheit; got != expectedFahrenheit {
		t.Errorf("ParaFahrenheit() = %.2f; esperado %.2f", got, expectedFahrenheit)
	}

	if got := temperature.Kelvin; got != expectedKelvin {
		t.Errorf("ParaKelvin() = %.2f; esperado %.2f", got, expectedKelvin)
	}
}
