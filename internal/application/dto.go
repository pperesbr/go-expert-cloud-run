package application

import "regexp"

type InputRequestDTO struct {
	ZipCode string `json:"zip_code"`
}

func (i *InputRequestDTO) Validate() bool {

	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(i.ZipCode)

}

type WeatherResponseDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}
