// Code generated by go-swagger; DO NOT EDIT.

package vp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppMessage AppMessage
//
// swagger:model AppMessage
type AppMessage struct {

	// attributes
	// Required: true
	Attributes []string `json:"attributes"`

	// Code
	// Required: true
	Code string `json:"code"`

	// Message
	// Required: true
	Message string `json:"message"`
}

// Validate validates this app message
func (m *AppMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppMessage) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	return nil
}

func (m *AppMessage) validateCode(formats strfmt.Registry) error {

	if err := validate.RequiredString("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *AppMessage) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this app message based on context it is used
func (m *AppMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppMessage) UnmarshalBinary(b []byte) error {
	var res AppMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
