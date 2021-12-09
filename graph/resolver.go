package graph

import (
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/endpoint"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DeviceRepository   device.Repository
	EndpointRepository endpoint.Repository
}
