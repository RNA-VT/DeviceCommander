package device

//Device represents a compliant physical component & its web address.
type Device struct {
	ID          string
	Name        string
	Description string
	Host        string
	Port        string
}

//NewDevice -
func NewDevice(host string, port string) (Device, error) {
	dev := Device{
		Host: host,
		Port: port,
	}
	//TODO: Load Device Data from db
	return dev, nil
}

//ToFullAddress returns a network address including the ip address and port that this micro is listening on
func (d Device) ToFullAddress() string {
	/* Just for pretty printing the micro info */
	return d.Host + ":" + d.Port
}
