// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineResponse2009 inline response 200 9
// swagger:model inline_response_200_9
type InlineResponse2009 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse2008EmbeddedLinks `json:"_links"`

	// allocation
	// Required: true
	Allocation []string `json:"allocation"`

	// aws instance id
	// Required: true
	AwsInstanceID *string `json:"aws_instance_id"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// docker name
	// Required: true
	DockerName *string `json:"docker_name"`

	// host
	// Required: true
	Host *string `json:"host"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// instance class
	// Required: true
	InstanceClass *string `json:"instance_class"`

	// layer
	// Required: true
	Layer *string `json:"layer"`

	// memory limit
	// Required: true
	MemoryLimit *int64 `json:"memory_limit"`

	// mounts
	// Required: true
	Mounts [][]string `json:"mounts"`

	// port
	// Required: true
	Port *int64 `json:"port"`

	// port mapping
	// Required: true
	PortMapping []int64 `json:"port_mapping"`

	// status
	// Required: true
	Status *string `json:"status"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this inline response 200 9
func (m *InlineResponse2009) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2009) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2009) validateAllocation(formats strfmt.Registry) error {

	if err := validate.Required("allocation", "body", m.Allocation); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateAwsInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("aws_instance_id", "body", m.AwsInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateDockerName(formats strfmt.Registry) error {

	if err := validate.Required("docker_name", "body", m.DockerName); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateInstanceClass(formats strfmt.Registry) error {

	if err := validate.Required("instance_class", "body", m.InstanceClass); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateLayer(formats strfmt.Registry) error {

	if err := validate.Required("layer", "body", m.Layer); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateMemoryLimit(formats strfmt.Registry) error {

	if err := validate.Required("memory_limit", "body", m.MemoryLimit); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateMounts(formats strfmt.Registry) error {

	if err := validate.Required("mounts", "body", m.Mounts); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validatePortMapping(formats strfmt.Registry) error {

	if err := validate.Required("port_mapping", "body", m.PortMapping); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2009) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2009) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2009) UnmarshalBinary(b []byte) error {
	var res InlineResponse2009
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
