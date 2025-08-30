package Hedayat

import (
	"errors"
)

type RegistrySchema interface {
	Schema() (string, error)
	Topic() (string, error)
}

type Boof struct {
	boofSchema string
}

func (b *Boof) Schema() (string, error) {
	return "", errors.New("ErrHedayad: boof schema is not defined")
}

func (b *Boof) Topic() (string, error) {
	return "", errors.New("ErrHedayad: boof topic is not defined")
}
