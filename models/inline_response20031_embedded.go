// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20031Embedded inline response 200 31 embedded
// swagger:model inline_response_200_31__embedded
type InlineResponse20031Embedded struct {

	// releases
	Releases []*InlineResponse20031EmbeddedReleases `json:"releases"`
}

// Validate validates this inline response 200 31 embedded
func (m *InlineResponse20031Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20031Embedded) validateReleases(formats strfmt.Registry) error {

	if swag.IsZero(m.Releases) { // not required
		return nil
	}

	for i := 0; i < len(m.Releases); i++ {
		if swag.IsZero(m.Releases[i]) { // not required
			continue
		}

		if m.Releases[i] != nil {
			if err := m.Releases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("releases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20031Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20031Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20031Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
