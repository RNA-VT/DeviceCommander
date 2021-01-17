package device

//Device represents a compliant physical component & its web address.
type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	failures    int
}

//NewDevice -
func NewDevice(host string, port int) (Device, error) {
	dev := Device{
		Host:     host,
		Port:     port,
		failures: 0,
	}
	//TODO: Load Device Data from db
	return dev, nil
}

//ToFullAddress returns a network address including the ip address and port that this micro is listening on
func (d Device) ToFullAddress() string {
	return d.Host + ":" + d.Port
}

//ProcessHealthCheckResult - updates health check failure count & returns
func (d *Device) ProcessHealthCheckResult(result bool) {
	if result { //Healthy
		d.failures = 0
	} else {
		d.failures++
	}
}

//Failed - If true, device should be deregistered
func (d Device) Failed() bool {
	failThreshold := 3
	return d.failures >= failThreshold
}
