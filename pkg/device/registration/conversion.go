package registration

import (
	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/scanner"
)

func newDeviceFromFoundDevice(d scanner.FoundDevice) device.NewDeviceParams {
	return device.NewDeviceParams{
		MAC:  &d.MAC,
		Host: d.IP,
		Port: d.Port,
	}
}
