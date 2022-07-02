package device

import "fmt"

func (d Device) RequestSpecification(client Client) (Device, error) {
	resp, err := client.Specification(d)
	if err != nil {
		d.logger.Warn(fmt.Sprintf("Error loading specification for [%s]", d.ID.String()))
		return d, err
	}

	spec, err := client.EvaluateSpecificationResponse(resp)
	if err != nil {
		d.logger.Warn(fmt.Sprintf("Error evaluating specification for [%s]", d.ID.String()))
		return d, err
	}

	return spec, nil
}

func (d Device) LoadFromSpecification(spec Device) Device {
	d.Description = spec.Description
	d.Endpoints = spec.Endpoints
	d.Name = spec.Name
	return d
}
