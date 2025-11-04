package data

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo)

type Data struct {
	// Database connection, Redis client, etc.
}

func NewData() (*Data, error) {
	return &Data{}, nil
}

