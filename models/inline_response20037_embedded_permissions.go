// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20037EmbeddedPermissions inline response 200 37 embedded permissions
// swagger:model inline_response_200_37__embedded_permissions
type InlineResponse20037EmbeddedPermissions struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse20037EmbeddedLinks `json:"_links,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this inline response 200 37 embedded permissions
func (m *InlineResponse20037EmbeddedPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20037EmbeddedPermissions) validateLinks(formats strfmt.Registry) error {

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
func (m *InlineResponse20037EmbeddedPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20037EmbeddedPermissions) UnmarshalBinary(b []byte) error {
	var res InlineResponse20037EmbeddedPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
