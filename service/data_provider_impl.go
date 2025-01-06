package service

import (
	"errors"
	"math/rand"
)

type DataProviderImpl struct {
}

// concrete implementation of the DataProvider interface
func (dp *DataProviderImpl) GetRandomData(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n must be greater than 0")
	}
	return rand.Intn(n+1), nil
}