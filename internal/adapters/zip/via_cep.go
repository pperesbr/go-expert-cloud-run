package zip

import (
	"encoding/json"
	"fmt"
	"github.com/pperesbr/go-expert-cloud-run/internal/application/ports"
	"io"
	"net/http"
	"strings"
)

type ViaCep struct {
}

func NewViaCep() *ViaCep {
	return &ViaCep{}
}

func (v *ViaCep) SearchZipInfo(zip string) (*ports.ZipInfo, error) {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		if resp.StatusCode == http.StatusBadRequest {
			return nil, ports.ErrZipInvalid
		}
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if _, ok := data["erro"]; ok {
		return nil, ports.ErrZipNotFound
	}

	var zipInfo ports.ZipInfo
	zipInfo.Code = strings.Replace(data["cep"].(string), "-", "", -1)
	zipInfo.City = data["localidade"].(string)
	zipInfo.UF = data["uf"].(string)
	zipInfo.State = data["estado"].(string)

	return &zipInfo, nil
}

var _ ports.ZipService = &ViaCep{}
