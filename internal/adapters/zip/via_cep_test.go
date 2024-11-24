package zip

import (
	"github.com/pperesbr/go-expert-cloud-run/internal/application/ports"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViaCep_SearchZipInfo(t *testing.T) {

	assert := assert.New(t)

	viaCep := NewViaCep()

	// Test invalid zip code
	_, err := viaCep.SearchZipInfo("aaaaaaa")
	assert.ErrorIs(err, ports.ErrZipInvalid)

	// Test not found zip code
	_, err = viaCep.SearchZipInfo("00000000")
	assert.ErrorIs(err, ports.ErrZipNotFound)

	// Test valid zip code
	zipInfo, err := viaCep.SearchZipInfo("87560000")
	assert.NoError(err)
	assert.Equal("87560000", zipInfo.Code)
	assert.Equal("Iporã", zipInfo.City)
	assert.Equal("PR", zipInfo.UF)
	assert.Equal("Paraná", zipInfo.State)

}
