package api

import (
	"encoding/json"
	"errors"
	"github.com/pperesbr/go-expert-cloud-run/internal/application"
	"github.com/pperesbr/go-expert-cloud-run/internal/application/ports"
	"net/http"
)

type Api struct {
	service *application.WeatherService
}

func NewApi(service *application.WeatherService) *Api {
	return &Api{service: service}
}

func (h *Api) GetWeather(w http.ResponseWriter, r *http.Request) {

	request := application.InputRequestDTO{ZipCode: r.PathValue("zipcode")}
	if !request.Validate() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid zip code"})
		return
	}

	response, err := h.service.GetWeather(request)

	if err != nil {
		if errors.Is(err, ports.ErrZipNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "can not find zipcode"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return

	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
