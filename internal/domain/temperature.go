package domain

import "math"

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func NewTemperature(celsius float64) *Temperature {

	scale := math.Pow(10, float64(1))

	return &Temperature{
		Celsius:    celsius,
		Fahrenheit: math.Round(celsius*1.8+32*scale) / scale,
		Kelvin:     math.Round(celsius+273.15*scale) / scale,
	}
}
