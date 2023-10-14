package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rna-vt/devicecommander/pkg/device/registration"
	"github.com/rna-vt/devicecommander/pkg/scanner"
	log "github.com/sirupsen/logrus"
)

type ARPScanResponse struct {
	Message string `json:"message"`
}

type OpsController struct {
	DeviceScanner   scanner.Scanner
	DeviceRegistrar registration.Registrar
}

func NewOpsController() (OpsController, error) {

	return OpsController{}, nil
}

func (controller OpsController) RunARPScan(c echo.Context) error {
	resp := ARPScanResponse{
		Message: "ARP Scan Started",
	}

	go func() {
		prospectiveDevices, err := controller.DeviceScanner.Scan()
		if err != nil {
			log.Errorf("error scanning for devices: %v", err)
			return
		}

		_, err = controller.DeviceRegistrar.HandleProspects(prospectiveDevices)
		if err != nil {
			log.Errorf("error handling prospective devices: %v", err)
			return
		}
	}()

	return c.JSON(http.StatusOK, resp)
}
