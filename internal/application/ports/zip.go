package ports

import "errors"

var ErrZipNotFound = errors.New("can not find zipcode")
var ErrZipInvalid = errors.New(" invalid zipcode")

type ZipInfo struct {
	Code  string
	City  string
	UF    string
	State string
}

func NewZipInfo(code, city, uf, state string) *ZipInfo {
	return &ZipInfo{
		Code:  code,
		City:  city,
		UF:    uf,
		State: state,
	}
}

type ZipService interface {
	SearchZipInfo(zip string) (*ZipInfo, error)
}
