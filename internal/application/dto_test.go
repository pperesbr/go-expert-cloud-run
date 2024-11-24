package application

import (
	"testing"
)

func TestInputRequestDTO_Validate(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  InputRequestDTO
		expected bool
	}{
		{
			name:     "CEP válido",
			zipCode:  InputRequestDTO{ZipCode: "12345678"},
			expected: true,
		},
		{
			name:     "CEP inválido",
			zipCode:  InputRequestDTO{ZipCode: "1234567"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.zipCode.Validate(); got != tt.expected {
				t.Errorf("ZipCode = %v; esperado %v", got, tt.expected)
			}
		})
	}

}
