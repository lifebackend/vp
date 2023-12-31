// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PingMessage PingMessage
//
// swagger:model PingMessage
type PingMessage struct {

	// ID
	// Required: true
	ID string `json:"id"`

	// password
	// Required: true
	Password string `json:"password"`
}

// Validate validates this ping message
func (m *PingMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PingMessage) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PingMessage) validatePassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ping message based on context it is used
func (m *PingMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PingMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingMessage) UnmarshalBinary(b []byte) error {
	var res PingMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
