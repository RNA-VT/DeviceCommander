package device

import (
	"github.com/rna-vt/devicecommander/src/utilities"
	"github.com/stretchr/testify/suite"
)

type SpecificationSuite struct {
	suite.Suite
}

func (s *SpecificationSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *SpecificationSuite) TestRequestSpecification() {

}
