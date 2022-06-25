package device

import "fmt"

func (d Device) RequestSpecification(client Client) error {
	resp, err := client.Specification(d)
	if err != nil {
		d.logger.Warn(fmt.Sprintf("Error loading specification for [%s]", d.ID.String()))
		return err
	}

	spec, err := client.EvaluateSpecificationResponse(resp)
	if err != nil {
		return err
	}

	d.LoadFromSpecification(spec)
	return nil
}

func (d *Device) LoadFromSpecification(spec Device) {
	d.Description = spec.Description
	d.Endpoints = spec.Endpoints
	d.Name = spec.Name
}
