package device

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/device/endpoint"
	"github.com/rna-vt/devicecommander/src/device/endpoint/parameter"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type SpecificationSuite struct {
	suite.Suite
}

func (s *SpecificationSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *SpecificationSuite) TestRequestSpecification() {
	device := Device{
		ID:          uuid.New(),
		Name:        "My Device",
		Description: "Devicey Devicey Deviceness Device",
		Endpoints: []endpoint.Endpoint{
			{
				Method:      "kabloom",
				Type:        "get",
				Description: new(string),
				Path:        new(string),
				Parameters:  []parameter.Parameter{},
			},
		},
	}
	spec := Device{
		Name:        "My Device",
		Description: "Devicey Devicey Deviceness Device",
		Endpoints: []endpoint.Endpoint{
			{
				Method:      "kabloom",
				Type:        "get",
				Description: new(string),
				Path:        new(string),
				Parameters:  []parameter.Parameter{},
			},
		},
	}

	// New, Healthy Device
	result, err := device.RequestSpecification(NewHTTPDeviceClient())
	s.Assertions.NoError(err)
	s.Assertions.Equal(spec, result)
}

func (s *SpecificationSuite) TestLoadFromSpecifcation() {
	device := Device{
		ID: uuid.New(),
	}
	spec := Device{
		Name:        "My Device",
		Description: "Devicey Devicey Deviceness Device",
		Endpoints: []endpoint.Endpoint{
			{
				Method:      "kabloom",
				Type:        "get",
				Description: new(string),
				Path:        new(string),
				Parameters:  []parameter.Parameter{},
			},
		},
	}

	device.LoadFromSpecification(spec)

	s.Assertions.Equal(spec.Name, device.Name)
	s.Assertions.Equal(spec.Description, device.Description)
	s.Assertions.Equal(spec.Endpoints, device.Endpoints)
}
