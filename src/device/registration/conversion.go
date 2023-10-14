package registration

import (
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/scanner"
)

func newDeviceFromFoundDevice(d scanner.FoundDevice) device.NewDeviceParams {
	return device.NewDeviceParams{
		MAC:  &d.MAC,
		Host: d.IP,
		Port: d.Port,
	}
}
