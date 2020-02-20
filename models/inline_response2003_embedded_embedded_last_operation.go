// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2003EmbeddedEmbeddedLastOperation inline response 200 3 embedded embedded last operation
// swagger:model inline_response_200_3__embedded__embedded_last_operation
type InlineResponse2003EmbeddedEmbeddedLastOperation struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse2003EmbeddedEmbeddedLastOperationLinks `json:"_links,omitempty"`

	// aborted
	Aborted bool `json:"aborted,omitempty"`

	// cancelled
	Cancelled bool `json:"cancelled,omitempty"`

	// certificate
	Certificate *string `json:"certificate,omitempty"`

	// command
	Command *string `json:"command,omitempty"`

	// container count
	ContainerCount *int64 `json:"container_count,omitempty"`

	// container size
	ContainerSize *int64 `json:"container_size,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// destination region
	DestinationRegion *string `json:"destination_region,omitempty"`

	// disk size
	DiskSize int64 `json:"disk_size,omitempty"`

	// docker ref
	DockerRef *string `json:"docker_ref,omitempty"`

	// env
	Env interface{} `json:"env,omitempty"`

	// git ref
	GitRef *string `json:"git_ref,omitempty"`

	// handle
	Handle *string `json:"handle,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// interactive
	Interactive *bool `json:"interactive,omitempty"`

	// private key
	PrivateKey *string `json:"private_key,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// user email
	UserEmail string `json:"user_email,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this inline response 200 3 embedded embedded last operation
func (m *InlineResponse2003EmbeddedEmbeddedLastOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2003EmbeddedEmbeddedLastOperation) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
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

// MarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedEmbeddedLastOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedEmbeddedLastOperation) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003EmbeddedEmbeddedLastOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
