package cluster

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/graph/model"
	"github.com/rna-vt/devicecommander/src/mocks"
	"github.com/rna-vt/devicecommander/src/test"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type ClusterSuite struct {
	suite.Suite
	mockDeviceService mocks.DeviceCRUDService
	mockDeviceClient  mocks.IDeviceClient
	cluster           Cluster
}

func (s *ClusterSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	s.mockDeviceService = mocks.DeviceCRUDService{}
	s.mockDeviceClient = mocks.IDeviceClient{}

	s.cluster = NewCluster(
		"testing",
		&s.mockDeviceService,
		&s.mockDeviceClient,
	)
}

func (s *ClusterSuite) GenerateDevices(count int) []*model.Device {
	devs := test.GenerateRandomNewDevices(count)
	collection := []*model.Device{}
	for _, d := range devs {
		tmpDev := device.DeviceFromNewDevice(d)
		collection = append(collection, &tmpDev)
	}
	return collection
}

func (s *ClusterSuite) TestRunHealthCheckLoop() {
	mockDevices := s.GenerateDevices(1)

	fmt.Println(len(mockDevices))

	s.mockDeviceService.On("Get", mock.AnythingOfType("model.Device")).Return(mockDevices, nil)

	tmpResponse := http.Response{
		Status: "200",
		Body:   io.NopCloser(strings.NewReader("healthy")),
	}

	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(&tmpResponse, nil)

	s.mockDeviceClient.On("EvaluateHealthCheckResponse", mock.AnythingOfType("*http.Response"), mock.AnythingOfType("device.Device")).Return(true)

	go s.cluster.RunHealthCheckLoop(1)

	time.Sleep(1 * time.Second)

	s.cluster.healthStop <- true

	s.mockDeviceService.AssertCalled(s.T(), "Get", model.Device{Active: true})

	s.mockDeviceService.AssertNumberOfCalls(s.T(), "Get", 1)

	// s.mockDeviceClient.AssertNumberOfCalls(s.T(), "Health", 3)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestClusterSuite(t *testing.T) {
	suite.Run(t, new(ClusterSuite))
}
