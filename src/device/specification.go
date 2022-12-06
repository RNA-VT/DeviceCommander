package device

func (d Device) LoadFromSpecification(spec *Device) Device {
	d.Description = spec.Description
	d.Endpoints = spec.Endpoints
	d.Name = spec.Name
	return d
}
